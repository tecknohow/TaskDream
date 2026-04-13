<template>
	<div class="eisenhower-view">
		<div class="eisenhower-header">
			<h1>Eisenhower Matrix</h1>
			<div class="header-actions">
				<QuickAddTask
					v-if="showCreateTask"
					@close="showCreateTask = false"
					@created="onTaskCreated"
				/>
				<button class="btn btn-primary" @click="showCreateTask = true">+ New Task</button>
			</div>
		</div>

		<div class="matrix-grid">
			<!-- Quadrant 1: Urgent & Important - DO -->
			<div class="quadrant q1">
				<div class="quadrant-header">
					<div class="quadrant-title">
						<span class="quadrant-icon">1</span>
						<div>
							<h2>Do First</h2>
							<span class="quadrant-desc">Urgent &amp; Important</span>
						</div>
					</div>
					<span class="task-count">{{ urgentImportant.length }}</span>
				</div>
				<div class="quadrant-tasks">
					<div
						v-for="task in urgentImportant"
						:key="task.id"
						class="quadrant-task-item"
						@click="navigateToTask(task)"
					>
						<span class="task-title">{{ task.title }}</span>
						<span v-if="task.dueDate" class="task-due">{{ formatDate(task.dueDate) }}</span>
					</div>
					<div v-if="urgentImportant.length === 0" class="empty-quadrant">
						No tasks here - great!
					</div>
				</div>
			</div>

			<!-- Quadrant 2: Not Urgent & Important - SCHEDULE -->
			<div class="quadrant q2">
				<div class="quadrant-header">
					<div class="quadrant-title">
						<span class="quadrant-icon">2</span>
						<div>
							<h2>Schedule</h2>
							<span class="quadrant-desc">Not Urgent &amp; Important</span>
						</div>
					</div>
					<span class="task-count">{{ notUrgentImportant.length }}</span>
				</div>
				<div class="quadrant-tasks">
					<div
						v-for="task in notUrgentImportant"
						:key="task.id"
						class="quadrant-task-item"
						@click="navigateToTask(task)"
					>
						<span class="task-title">{{ task.title }}</span>
						<span v-if="task.dueDate" class="task-due">{{ formatDate(task.dueDate) }}</span>
					</div>
					<div v-if="notUrgentImportant.length === 0" class="empty-quadrant">
						Plan important work here
					</div>
				</div>
			</div>

			<!-- Quadrant 3: Urgent & Not Important - DELEGATE -->
			<div class="quadrant q3">
				<div class="quadrant-header">
					<div class="quadrant-title">
						<span class="quadrant-icon">3</span>
						<div>
							<h2>Delegate</h2>
							<span class="quadrant-desc">Urgent &amp; Not Important</span>
						</div>
					</div>
					<span class="task-count">{{ urgentNotImportant.length }}</span>
				</div>
				<div class="quadrant-tasks">
					<div
						v-for="task in urgentNotImportant"
						:key="task.id"
						class="quadrant-task-item"
						@click="navigateToTask(task)"
					>
						<span class="task-title">{{ task.title }}</span>
						<span v-if="task.dueDate" class="task-due">{{ formatDate(task.dueDate) }}</span>
					</div>
					<div v-if="urgentNotImportant.length === 0" class="empty-quadrant">
						Delegate these tasks
					</div>
				</div>
			</div>

			<!-- Quadrant 4: Not Urgent & Not Important - ELIMINATE -->
			<div class="quadrant q4">
				<div class="quadrant-header">
					<div class="quadrant-title">
						<span class="quadrant-icon">4</span>
						<div>
							<h2>Eliminate</h2>
							<span class="quadrant-desc">Not Urgent &amp; Not Important</span>
						</div>
					</div>
					<span class="task-count">{{ notUrgentNotImportant.length }}</span>
				</div>
				<div class="quadrant-tasks">
					<div
						v-for="task in notUrgentNotImportant"
						:key="task.id"
						class="quadrant-task-item"
						@click="navigateToTask(task)"
					>
						<span class="task-title">{{ task.title }}</span>
						<span v-if="task.dueDate" class="task-due">{{ formatDate(task.dueDate) }}</span>
					</div>
					<div v-if="notUrgentNotImportant.length === 0" class="empty-quadrant">
						Consider removing these
					</div>
				</div>
			</div>
		</div>

		<!-- Axis labels -->
		<div class="axis-labels">
			<div class="axis-y-label">
				<span class="axis-arrow">&uarr;</span> IMPORTANT
			</div>
			<div class="axis-x-label">
				URGENT <span class="axis-arrow">&rarr;</span>
			</div>
		</div>
	</div>
</template>

<script lang="ts" setup>
import {ref, computed, onMounted} from 'vue'
import {useRouter} from 'vue-router'

import QuickAddTask from '@/components/QuickAddTask.vue'
import TaskService from '@/services/task'
import TaskModel from '@/models/task'
import type {ITask} from '@/modelTypes/ITask'
import {PRIORITIES} from '@/constants/priorities'

const router = useRouter()
const taskService = new TaskService()

const allTasks = ref<ITask[]>([])
const showCreateTask = ref(false)

// Eisenhower quadrant classification based on priority and due date:
// - Urgent: has a due date within the next 2 days, or priority >= URGENT
// - Important: priority >= HIGH
const isUrgent = (task: ITask): boolean => {
	if (task.priority >= PRIORITIES.URGENT) return true
	if (task.dueDate) {
		const due = new Date(task.dueDate)
		const twoDaysFromNow = new Date()
		twoDaysFromNow.setDate(twoDaysFromNow.getDate() + 2)
		return due <= twoDaysFromNow
	}
	return false
}

const isImportant = (task: ITask): boolean => {
	return task.priority >= PRIORITIES.HIGH
}

const openTasks = computed(() => allTasks.value.filter(t => !t.done))

const urgentImportant = computed(() =>
	openTasks.value.filter(t => isUrgent(t) && isImportant(t)),
)
const notUrgentImportant = computed(() =>
	openTasks.value.filter(t => !isUrgent(t) && isImportant(t)),
)
const urgentNotImportant = computed(() =>
	openTasks.value.filter(t => isUrgent(t) && !isImportant(t)),
)
const notUrgentNotImportant = computed(() =>
	openTasks.value.filter(t => !isUrgent(t) && !isImportant(t)),
)

onMounted(async () => {
	try {
		allTasks.value = await taskService.getAll(new TaskModel({}), {
			filter: 'done = false',
		})
	} catch {
		// Loading failed
	}
})

function navigateToTask(task: ITask) {
	router.push({name: 'task.detail', params: {id: task.id}})
}

function onTaskCreated() {
	showCreateTask.value = false
	// Reload tasks
	taskService.getAll(new TaskModel({}), {filter: 'done = false'}).then(tasks => {
		allTasks.value = tasks
	})
}

function formatDate(date: Date | null): string {
	if (!date) return ''
	return new Date(date).toLocaleDateString('en-US', {
		month: 'short',
		day: 'numeric',
	})
}
</script>

<style scoped lang="scss">
.eisenhower-view {
	max-width: 1400px;
	margin: 0 auto;
	position: relative;
}

.eisenhower-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: 1rem;

	h1 { margin: 0; }
}

.matrix-grid {
	display: grid;
	grid-template-columns: 1fr 1fr;
	grid-template-rows: 1fr 1fr;
	gap: .75rem;
	min-height: 600px;
}

.quadrant {
	background-color: var(--white);
	border-radius: $radius;
	display: flex;
	flex-direction: column;
	overflow: hidden;
	box-shadow: var(--shadow-sm);
	border: 2px solid transparent;

	&.q1 {
		border-color: rgba(239, 68, 68, 0.3);
		.quadrant-icon { background: rgba(239, 68, 68, 0.15); color: var(--danger); }
	}

	&.q2 {
		border-color: rgba(37, 99, 235, 0.3);
		.quadrant-icon { background: rgba(37, 99, 235, 0.15); color: var(--primary); }
	}

	&.q3 {
		border-color: rgba(245, 158, 11, 0.3);
		.quadrant-icon { background: rgba(245, 158, 11, 0.15); color: var(--warning); }
	}

	&.q4 {
		border-color: rgba(156, 163, 175, 0.3);
		.quadrant-icon { background: rgba(156, 163, 175, 0.15); color: var(--grey-500); }
	}
}

.quadrant-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: .75rem 1rem;
	border-bottom: 1px solid var(--grey-200);
}

.quadrant-title {
	display: flex;
	align-items: center;
	gap: .75rem;

	h2 {
		margin: 0;
		font-size: 1rem;
	}
}

.quadrant-icon {
	width: 32px;
	height: 32px;
	border-radius: 50%;
	display: flex;
	align-items: center;
	justify-content: center;
	font-weight: 700;
	font-size: .875rem;
}

.quadrant-desc {
	font-size: .75rem;
	color: var(--grey-500);
}

.task-count {
	background-color: var(--grey-100);
	color: var(--grey-500);
	padding: .25rem .5rem;
	border-radius: $radius;
	font-size: .75rem;
	font-weight: 600;
}

.quadrant-tasks {
	flex: 1;
	padding: .75rem;
	overflow-y: auto;
	display: flex;
	flex-direction: column;
	gap: .5rem;
}

.quadrant-task-item {
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: .5rem .75rem;
	border: 1px solid var(--grey-200);
	border-radius: $radius;
	cursor: pointer;
	font-size: .875rem;
	transition: all $transition;

	&:hover {
		border-color: var(--primary);
		background: var(--grey-50);
	}

	.task-title {
		font-weight: 500;
	}

	.task-due {
		font-size: .75rem;
		color: var(--grey-500);
	}
}

.empty-quadrant {
	display: flex;
	align-items: center;
	justify-content: center;
	height: 100%;
	min-height: 100px;
	color: var(--grey-400);
	font-size: .875rem;
}

.axis-labels {
	position: relative;
	margin-top: .75rem;
}

.axis-y-label, .axis-x-label {
	font-size: .75rem;
	color: var(--grey-400);
	text-transform: uppercase;
	letter-spacing: 1px;
	font-weight: 600;
}

.axis-y-label {
	position: absolute;
	left: -10px;
	top: -320px;
	transform: rotate(-90deg);
	transform-origin: left center;
}

.axis-x-label {
	text-align: right;
}

@media (max-width: 768px) {
	.matrix-grid {
		grid-template-columns: 1fr;
		min-height: auto;
	}

	.axis-y-label {
		display: none;
	}
}
</style>
