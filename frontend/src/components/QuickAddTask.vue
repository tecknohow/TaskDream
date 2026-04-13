<template>
	<Teleport to="body">
		<div class="quick-add-overlay" @click.self="$emit('close')" @keydown.esc="$emit('close')">
			<div class="quick-add-modal" ref="modalRef">
				<div class="quick-add-header">
					<span class="quick-add-icon">+</span>
					<span>Quick Add Task</span>
					<kbd class="kbd">Esc</kbd>
				</div>

				<form @submit.prevent="submit" class="quick-add-form">
					<input
						ref="titleInput"
						v-model="title"
						type="text"
						placeholder="What needs to be done?"
						class="quick-add-input"
						autofocus
					/>

					<div class="quick-add-options">
						<div class="priority-selector">
							<button
								v-for="p in priorities"
								:key="p.value"
								type="button"
								class="priority-btn"
								:class="{ active: priority === p.value, [`priority-${p.value}`]: true }"
								@click="priority = p.value"
								:title="p.label"
							>
								{{ p.icon }}
							</button>
						</div>

						<input
							v-model="dueDate"
							type="date"
							class="option-date"
							placeholder="Due date"
						/>
					</div>

					<div class="quick-add-footer">
						<div class="quick-add-actions">
							<button type="button" class="btn btn-ghost" @click="$emit('close')">Cancel</button>
							<button type="submit" class="btn btn-primary" :disabled="!title.trim() || loading">
								{{ loading ? 'Adding...' : 'Add Task' }}
							</button>
						</div>
					</div>
				</form>
			</div>
		</div>
	</Teleport>
</template>

<script lang="ts" setup>
import {ref, onMounted} from 'vue'

import TaskService from '@/services/task'
import TaskModel from '@/models/task'
import type {ITask} from '@/modelTypes/ITask'
import {PRIORITIES} from '@/constants/priorities'
import {useAuthStore} from '@/stores/auth'

const emit = defineEmits<{
	close: []
	created: [task: ITask]
}>()

const authStore = useAuthStore()
const taskService = new TaskService()

const titleInput = ref<HTMLInputElement>()
const title = ref('')
const priority = ref(PRIORITIES.UNSET)
const dueDate = ref('')
const loading = ref(false)

const priorities = [
	{value: PRIORITIES.UNSET, label: 'None', icon: '-'},
	{value: PRIORITIES.LOW, label: 'Low', icon: '!'},
	{value: PRIORITIES.MEDIUM, label: 'Medium', icon: '!!'},
	{value: PRIORITIES.HIGH, label: 'High', icon: '!!!'},
	{value: PRIORITIES.URGENT, label: 'Urgent', icon: '!!!!'},
]

onMounted(() => {
	titleInput.value?.focus()
})

async function submit() {
	if (!title.value.trim()) return

	loading.value = true
	try {
		const projectId = authStore.settings?.defaultProjectId || 1

		const newTask = new TaskModel({
			title: title.value.trim(),
			projectId,
			priority: priority.value,
			dueDate: dueDate.value ? new Date(dueDate.value) : null,
		})

		const created = await taskService.create(newTask)
		emit('created', created)

		// Reset for another quick add
		title.value = ''
		dueDate.value = ''
		priority.value = PRIORITIES.UNSET
		titleInput.value?.focus()
	} catch {
		// Error handled
	} finally {
		loading.value = false
	}
}
</script>

<style scoped lang="scss">
.quick-add-overlay {
	position: fixed;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	background-color: rgba(0, 0, 0, 0.5);
	display: flex;
	align-items: flex-start;
	justify-content: center;
	padding-top: 15vh;
	z-index: 100;
	backdrop-filter: blur(4px);
}

.quick-add-modal {
	background-color: var(--white);
	border-radius: $radius;
	box-shadow: var(--shadow-lg);
	width: 560px;
	max-width: 90vw;
	overflow: hidden;
	animation: slideIn 0.15s ease-out;
}

.quick-add-header {
	display: flex;
	align-items: center;
	gap: .5rem;
	padding: .75rem 1rem;
	border-bottom: 1px solid var(--grey-200);
	font-size: .875rem;
	color: var(--grey-500);
}

.quick-add-icon {
	font-size: 1.125rem;
	font-weight: 700;
	color: var(--primary);
}

.kbd {
	margin-left: auto;
	padding: 2px 6px;
	background-color: var(--grey-100);
	border: 1px solid var(--grey-200);
	border-radius: $radius;
	font-size: .75rem;
	font-family: monospace;
}

.quick-add-form {
	padding: 1rem;
}

.quick-add-input {
	width: 100%;
	border: none;
	font-size: 1.125rem;
	padding: .5rem 0;
	background: transparent;
	color: var(--text);

	&:focus {
		outline: none;
		box-shadow: none;
	}

	&::placeholder {
		color: var(--grey-400);
	}
}

.quick-add-options {
	display: flex;
	gap: .75rem;
	margin-top: .75rem;
	flex-wrap: wrap;
}

.option-date {
	padding: .25rem .5rem;
	border-radius: $radius;
	font-size: .875rem;
}

.priority-selector {
	display: flex;
	gap: 2px;
	background-color: var(--grey-100);
	border-radius: $radius;
	padding: 2px;
}

.priority-btn {
	padding: .25rem .5rem;
	border: none;
	background: transparent;
	border-radius: $radius;
	cursor: pointer;
	font-size: .75rem;
	font-weight: 700;
	color: var(--grey-400);
	transition: all $transition;

	&.active {
		background-color: var(--white);
		box-shadow: var(--shadow-sm);
	}

	&.active.priority-0 { color: var(--grey-500); }
	&.active.priority-1 { color: var(--success); }
	&.active.priority-2 { color: var(--warning); }
	&.active.priority-3 { color: var(--danger); }
	&.active.priority-4 { color: var(--danger); font-weight: 900; }
}

.quick-add-footer {
	display: flex;
	justify-content: flex-end;
	align-items: center;
	margin-top: 1rem;
	padding-top: .75rem;
	border-top: 1px solid var(--grey-200);
}

.quick-add-actions {
	display: flex;
	gap: .5rem;
}

@keyframes slideIn {
	from {
		opacity: 0;
		transform: translateY(-10px);
	}
	to {
		opacity: 1;
		transform: translateY(0);
	}
}
</style>
