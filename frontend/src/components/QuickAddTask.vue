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
            <select v-model="projectId" class="option-select">
              <option :value="0">No Project</option>
              <option
                v-for="project in projectsStore.activeProjects"
                :key="project.id"
                :value="project.id"
              >
                {{ project.title }}
              </option>
            </select>

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
            <div class="eisenhower-toggles">
              <label class="toggle-label">
                <input type="checkbox" v-model="isUrgent" />
                <span>Urgent</span>
              </label>
              <label class="toggle-label">
                <input type="checkbox" v-model="isImportant" />
                <span>Important</span>
              </label>
            </div>
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

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useTasksStore } from '@/stores/tasks'
import { useProjectsStore } from '@/stores/projects'

const emit = defineEmits(['close', 'created'])

const tasksStore = useTasksStore()
const projectsStore = useProjectsStore()

const titleInput = ref<HTMLInputElement>()
const title = ref('')
const projectId = ref(0)
const priority = ref(0)
const dueDate = ref('')
const isUrgent = ref(false)
const isImportant = ref(false)
const loading = ref(false)

const priorities = [
  { value: 0, label: 'Low', icon: '!' },
  { value: 1, label: 'Medium', icon: '!!' },
  { value: 2, label: 'High', icon: '!!!' },
  { value: 3, label: 'Urgent', icon: '!!!!' }
]

onMounted(() => {
  titleInput.value?.focus()
})

const submit = async () => {
  if (!title.value.trim()) return

  loading.value = true
  try {
    const task = await tasksStore.create({
      title: title.value.trim(),
      project_id: projectId.value || undefined,
      priority: priority.value,
      urgency: isUrgent.value ? 1 : 0,
      importance: isImportant.value ? 1 : 0,
      due_date: dueDate.value || undefined
    } as any)

    emit('created', task)
    // Reset for another quick add
    title.value = ''
    dueDate.value = ''
    titleInput.value?.focus()
  } catch {
    // error handled in store
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
  background-color: var(--bg-primary);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-xl);
  width: 560px;
  max-width: 90vw;
  overflow: hidden;
  animation: slideIn 0.15s ease-out;
}

.quick-add-header {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-md) var(--spacing-lg);
  border-bottom: 1px solid var(--border-color);
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
}

.quick-add-icon {
  font-size: var(--font-size-lg);
  font-weight: 700;
  color: var(--color-primary);
}

.kbd {
  margin-left: auto;
  padding: 2px 6px;
  background-color: var(--bg-tertiary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-family: monospace;
}

.quick-add-form {
  padding: var(--spacing-lg);
}

.quick-add-input {
  width: 100%;
  border: none;
  font-size: var(--font-size-lg);
  padding: var(--spacing-sm) 0;
  background: transparent;
  color: var(--text-primary);

  &:focus {
    outline: none;
    box-shadow: none;
  }

  &::placeholder {
    color: var(--text-tertiary);
  }
}

.quick-add-options {
  display: flex;
  gap: var(--spacing-md);
  margin-top: var(--spacing-md);
  flex-wrap: wrap;
}

.option-select {
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-md);
  font-size: var(--font-size-sm);
  min-width: 140px;
}

.option-date {
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-md);
  font-size: var(--font-size-sm);
}

.priority-selector {
  display: flex;
  gap: 2px;
  background-color: var(--bg-tertiary);
  border-radius: var(--radius-md);
  padding: 2px;
}

.priority-btn {
  padding: var(--spacing-xs) var(--spacing-sm);
  border: none;
  background: transparent;
  border-radius: var(--radius-sm);
  cursor: pointer;
  font-size: var(--font-size-xs);
  font-weight: 700;
  color: var(--text-tertiary);
  transition: all var(--transition-fast);

  &.active {
    background-color: var(--bg-primary);
    box-shadow: var(--shadow-sm);
  }

  &.active.priority-0 { color: var(--color-success); }
  &.active.priority-1 { color: var(--color-warning); }
  &.active.priority-2 { color: #dc2626; }
  &.active.priority-3 { color: var(--color-error); }
}

.quick-add-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: var(--spacing-lg);
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--border-color);
}

.eisenhower-toggles {
  display: flex;
  gap: var(--spacing-md);
}

.toggle-label {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
  cursor: pointer;

  input {
    cursor: pointer;
  }
}

.quick-add-actions {
  display: flex;
  gap: var(--spacing-sm);
}
</style>
