<template>
	<div class="focus-mode" :class="{ 'is-active': pomodoroStore.isRunning || pomodoroStore.isOnBreak }">
		<div class="focus-container">
			<!-- Task selector when no task selected -->
			<div v-if="!selectedTask" class="task-selector">
				<h1>Focus Mode</h1>
				<p class="text-secondary">Select a task to start a focused work session</p>

				<div class="task-search">
					<input
						v-model="searchQuery"
						type="text"
						placeholder="Search tasks..."
						class="search-input"
					/>
				</div>

				<div class="task-list">
					<div
						v-for="task in filteredTasks"
						:key="task.id"
						class="focus-task-item"
						@click="selectTask(task)"
					>
						<div class="task-info">
							<div class="task-title">{{ task.title }}</div>
							<div class="task-meta-row">
								<span v-if="task.priority" class="meta-badge" :class="`priority-${task.priority}`">
									{{ priorityLabel(task.priority) }}
								</span>
								<span v-if="task.percentDone > 0" class="meta-badge">
									{{ task.percentDone }}%
								</span>
							</div>
						</div>
						<button class="btn btn-primary btn-sm">Focus</button>
					</div>

					<div v-if="filteredTasks.length === 0" class="empty-state">
						{{ searchQuery ? 'No matching tasks' : 'No open tasks available' }}
					</div>
				</div>
			</div>

			<!-- Active focus session -->
			<div v-else class="focus-session">
				<button class="btn btn-ghost back-btn" @click="deselectTask">
					&larr; Back to task list
				</button>

				<div class="focus-task-header">
					<h2>{{ selectedTask.title }}</h2>
					<p v-if="selectedTask.description" class="task-description">
						{{ selectedTask.description }}
					</p>
				</div>

				<PomodoroTimer
					:task-id="selectedTask.id"
					@start="startFocus"
				/>

				<!-- Related subtasks -->
				<div v-if="subtasks.length > 0" class="subtasks-section">
					<h3>Subtasks</h3>
					<div class="subtask-list">
						<div v-for="st in subtasks" :key="st.id" class="subtask-item">
							<input
								type="checkbox"
								:checked="st.done"
								@change="toggleSubtask(st)"
							/>
							<span :class="{ done: st.done }">{{ st.title }}</span>
						</div>
					</div>
				</div>

				<!-- Task stats panel -->
				<div class="focus-stats-panel">
					<div class="stat-card">
						<div class="stat-label">Priority</div>
						<div class="stat-value">
							{{ priorityLabel(selectedTask.priority) }}
						</div>
					</div>
					<div class="stat-card">
						<div class="stat-label">Due Date</div>
						<div class="stat-value">
							{{ selectedTask.dueDate ? formatDate(selectedTask.dueDate) : 'N/A' }}
						</div>
					</div>
					<div class="stat-card">
						<div class="stat-label">Progress</div>
						<div class="stat-value">{{ selectedTask.percentDone }}%</div>
					</div>
				</div>

				<div class="focus-actions">
					<button v-if="!selectedTask.done" class="btn btn-primary" @click="markDone">
						Mark as Done
					</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts" setup>
import {ref, computed, onMounted} from 'vue'

import PomodoroTimer from '@/components/PomodoroTimer.vue'
import {usePomodoroStore} from '@/stores/pomodoro'

import TaskService from '@/services/task'
import TaskModel from '@/models/task'
import type {ITask} from '@/modelTypes/ITask'
import {PRIORITIES} from '@/constants/priorities'

const pomodoroStore = usePomodoroStore()
const taskService = new TaskService()

const allTasks = ref<ITask[]>([])
const selectedTask = ref<ITask | null>(null)
const subtasks = ref<ITask[]>([])
const searchQuery = ref('')

const filteredTasks = computed(() => {
	const tasks = allTasks.value.filter(t => !t.done)
	if (!searchQuery.value) return tasks
	const q = searchQuery.value.toLowerCase()
	return tasks.filter(t =>
		t.title.toLowerCase().includes(q) ||
		(t.description && t.description.toLowerCase().includes(q)),
	)
})

onMounted(async () => {
	await pomodoroStore.fetchSettings()
	await pomodoroStore.fetchStats()
	try {
		allTasks.value = await taskService.getAll(new TaskModel({}), {
			filter: 'done = false',
			sort_by: ['priority', 'due_date'],
			order_by: ['desc', 'asc'],
		})
	} catch {
		// Tasks loading failed
	}
})

function priorityLabel(priority: number): string {
	const labels: Record<number, string> = {
		[PRIORITIES.UNSET]: 'None',
		[PRIORITIES.LOW]: 'Low',
		[PRIORITIES.MEDIUM]: 'Medium',
		[PRIORITIES.HIGH]: 'High',
		[PRIORITIES.URGENT]: 'Urgent',
		[PRIORITIES.DO_NOW]: 'DO NOW',
	}
	return labels[priority] || 'None'
}

function formatDate(date: Date | null): string {
	if (!date) return 'N/A'
	return new Date(date).toLocaleDateString('en-US', {
		month: 'short',
		day: 'numeric',
		year: 'numeric',
	})
}

async function selectTask(task: ITask) {
	selectedTask.value = task
	// Try to get related subtask info
	try {
		const fullTask = await taskService.get(new TaskModel({id: task.id}))
		if (fullTask.relatedTasks) {
			const subtaskRelation = fullTask.relatedTasks.subtask
			if (subtaskRelation && Array.isArray(subtaskRelation)) {
				subtasks.value = subtaskRelation
			}
		}
	} catch {
		subtasks.value = []
	}
}

function deselectTask() {
	if (pomodoroStore.isRunning) {
		pomodoroStore.cancelPomodoro()
	}
	selectedTask.value = null
	subtasks.value = []
}

function startFocus() {
	if (selectedTask.value) {
		pomodoroStore.startPomodoro(selectedTask.value.id)
	}
}

async function toggleSubtask(st: ITask) {
	try {
		await taskService.update(new TaskModel({...st, done: !st.done}))
		st.done = !st.done
	} catch {
		// Toggle failed
	}
}

async function markDone() {
	if (selectedTask.value) {
		try {
			await taskService.update(new TaskModel({...selectedTask.value, done: true}))
			selectedTask.value = null
		} catch {
			// Mark done failed
		}
	}
}
</script>

<style scoped lang="scss">
.focus-mode {
	max-width: 700px;
	margin: 0 auto;
	min-height: 80vh;

	&.is-active {
		.focus-container {
			background: linear-gradient(135deg, var(--white) 0%, var(--grey-100) 100%);
		}
	}
}

.focus-container {
	background-color: var(--white);
	border-radius: $radius;
	padding: 2rem;
	box-shadow: var(--shadow-sm);
}

.task-selector {
	h1 {
		text-align: center;
		margin-bottom: .25rem;
	}

	p {
		text-align: center;
		margin-bottom: 2rem;
	}
}

.search-input {
	width: 100%;
	padding: .75rem;
	border: 1px solid var(--grey-300);
	border-radius: $radius;
	font-size: 1rem;
	margin-bottom: 1rem;
}

.task-list {
	display: flex;
	flex-direction: column;
	gap: .5rem;
	max-height: 500px;
	overflow-y: auto;
}

.focus-task-item {
	display: flex;
	align-items: center;
	justify-content: space-between;
	padding: .75rem;
	border: 1px solid var(--grey-300);
	border-radius: $radius;
	cursor: pointer;
	transition: all $transition;

	&:hover {
		border-color: var(--primary);
		box-shadow: var(--shadow-sm);
	}
}

.task-info {
	flex: 1;
}

.task-title {
	font-weight: 500;
	margin-bottom: 2px;
}

.task-meta-row {
	display: flex;
	gap: .5rem;
}

.meta-badge {
	font-size: .75rem;
	color: var(--grey-500);
	padding: 1px 4px;
	background: var(--grey-100);
	border-radius: $radius;

	&.priority-3 { color: var(--warning); }
	&.priority-4 { color: var(--danger); }
	&.priority-5 { color: var(--danger); font-weight: 700; }
}

.back-btn {
	margin-bottom: 1rem;
}

.focus-task-header {
	text-align: center;
	margin-bottom: 1rem;

	h2 { margin-bottom: .5rem; }

	.task-description {
		color: var(--grey-500);
		font-size: .875rem;
	}
}

.subtasks-section {
	margin-top: 2rem;
	padding-top: 1rem;
	border-top: 1px solid var(--grey-200);

	h3 {
		margin-bottom: .75rem;
		font-size: 1rem;
	}
}

.subtask-list {
	display: flex;
	flex-direction: column;
	gap: .5rem;
}

.subtask-item {
	display: flex;
	align-items: center;
	gap: .5rem;
	padding: .5rem;
	font-size: .875rem;

	.done {
		text-decoration: line-through;
		color: var(--grey-400);
	}
}

.focus-stats-panel {
	display: grid;
	grid-template-columns: repeat(3, 1fr);
	gap: .75rem;
	margin-top: 2rem;
	padding-top: 1rem;
	border-top: 1px solid var(--grey-200);
}

.stat-card {
	text-align: center;
	padding: .75rem;
	background: var(--grey-100);
	border-radius: $radius;
}

.stat-label {
	font-size: .75rem;
	color: var(--grey-500);
	text-transform: uppercase;
	margin-bottom: .25rem;
}

.stat-value {
	font-size: 1.125rem;
	font-weight: 700;
	color: var(--primary);
}

.focus-actions {
	display: flex;
	justify-content: center;
	margin-top: 1rem;
}

.empty-state {
	text-align: center;
	padding: 2rem;
	color: var(--grey-400);
}
</style>
