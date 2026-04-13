<template>
  <div class="modal-overlay" @click.self="close">
    <div class="modal">
      <div class="modal-header">
        <h2>Create New Task</h2>
        <button class="btn btn-ghost" @click="close">✕</button>
      </div>

      <form @submit.prevent="createTask" class="modal-form">
        <div class="form-group">
          <label for="title">Task Title *</label>
          <input
            id="title"
            v-model="form.title"
            type="text"
            placeholder="What needs to be done?"
            required
          />
        </div>

        <div class="form-group">
          <label for="description">Description</label>
          <textarea
            id="description"
            v-model="form.description"
            placeholder="Add more details..."
          ></textarea>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label for="priority">Priority</label>
            <select id="priority" v-model.number="form.priority">
              <option value="0">Low</option>
              <option value="1">Medium</option>
              <option value="2">High</option>
              <option value="3">Urgent</option>
            </select>
          </div>

          <div class="form-group">
            <label for="dueDate">Due Date</label>
            <input
              id="dueDate"
              v-model="form.dueDate"
              type="date"
            />
          </div>
        </div>

        <div class="form-actions">
          <button type="button" class="btn btn-secondary" @click="close">
            Cancel
          </button>
          <button type="submit" class="btn btn-primary" :disabled="loading">
            {{ loading ? 'Creating...' : 'Create Task' }}
          </button>
        </div>

        <div v-if="error" class="error-message">
          {{ error }}
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useTasksStore } from '@/stores/tasks'

defineProps<{
  projectId?: string
}>()

const emit = defineEmits<{
  close: []
}>()

const tasksStore = useTasksStore()

const form = ref({
  title: '',
  description: '',
  priority: 1,
  dueDate: ''
})

const loading = ref(false)
const error = ref('')

const createTask = async () => {
  loading.value = true
  error.value = ''

  try {
    await tasksStore.create({
      title: form.value.title,
      description: form.value.description,
      priority: form.value.priority,
      dueDate: form.value.dueDate ? new Date(form.value.dueDate) : undefined,
      status: 'todo',
      projectId: 'default',
      createdAt: new Date(),
      updatedAt: new Date()
    })

    close()
  } catch (err: any) {
    error.value = err || 'Failed to create task'
  } finally {
    loading.value = false
  }
}

const close = () => {
  form.value = { title: '', description: '', priority: 1, dueDate: '' }
  emit('close')
}
</script>

<style scoped lang="scss">
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 50;
}

.modal {
  background-color: var(--bg-primary);
  border-radius: var(--radius-lg);
  padding: var(--spacing-xl);
  box-shadow: var(--shadow-xl);
  max-width: 500px;
  width: 90%;
  z-index: 51;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);

  h2 {
    margin: 0;
  }
}

.modal-form {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);

  label {
    font-weight: 500;
    font-size: var(--font-size-sm);
  }

  input,
  textarea,
  select {
    padding: var(--spacing-md);
    border: 1px solid var(--border-color);
    border-radius: var(--radius-md);
    font-size: var(--font-size-base);
    background-color: var(--bg-secondary);
    color: var(--text-primary);

    &:focus {
      outline: none;
      border-color: var(--color-primary);
      background-color: var(--bg-primary);
    }
  }

  textarea {
    resize: vertical;
    min-height: 100px;
  }
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-md);
}

.form-actions {
  display: flex;
  gap: var(--spacing-md);
  justify-content: flex-end;
  margin-top: var(--spacing-md);
  border-top: 1px solid var(--border-color);
  padding-top: var(--spacing-md);
}

.error-message {
  padding: var(--spacing-md);
  background-color: rgba(239, 68, 68, 0.1);
  color: var(--color-error);
  border-radius: var(--radius-md);
  font-size: var(--font-size-sm);
}
</style>
