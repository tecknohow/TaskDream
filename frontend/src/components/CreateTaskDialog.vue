<template>
  <Teleport to="body">
    <div class="modal-overlay" @click.self="$emit('close')">
      <div class="modal">
        <h2>Create Task</h2>

        <form @submit.prevent="submit">
          <div class="form-group">
            <label>Title</label>
            <input v-model="title" type="text" placeholder="Task title" required autofocus />
          </div>

          <div class="form-group">
            <label>Description</label>
            <textarea v-model="description" placeholder="Task description" rows="3"></textarea>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>Project</label>
              <select v-model.number="projectId">
                <option :value="0">No Project</option>
                <option
                  v-for="project in projectsStore.activeProjects"
                  :key="project.id"
                  :value="project.id"
                >
                  {{ project.title }}
                </option>
              </select>
            </div>

            <div class="form-group">
              <label>Priority</label>
              <select v-model.number="priority">
                <option :value="0">Low</option>
                <option :value="1">Medium</option>
                <option :value="2">High</option>
                <option :value="3">Urgent</option>
              </select>
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label>Due Date</label>
              <input v-model="dueDate" type="date" />
            </div>

            <div class="form-group">
              <label>Estimated Time (minutes)</label>
              <input v-model.number="estimatedMinutes" type="number" min="0" placeholder="e.g. 30" />
            </div>
          </div>

          <div class="form-row">
            <div class="form-group">
              <label class="toggle-label">
                <input type="checkbox" v-model="isUrgent" />
                Urgent (Eisenhower)
              </label>
            </div>
            <div class="form-group">
              <label class="toggle-label">
                <input type="checkbox" v-model="isImportant" />
                Important (Eisenhower)
              </label>
            </div>
          </div>

          <div class="form-actions">
            <button type="button" class="btn btn-ghost" @click="$emit('close')">Cancel</button>
            <button type="submit" class="btn btn-primary" :disabled="!title.trim() || loading">
              {{ loading ? 'Creating...' : 'Create Task' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useTasksStore } from '@/stores/tasks'
import { useProjectsStore } from '@/stores/projects'

const emit = defineEmits(['close', 'created'])

const tasksStore = useTasksStore()
const projectsStore = useProjectsStore()

const title = ref('')
const description = ref('')
const projectId = ref(0)
const priority = ref(0)
const dueDate = ref('')
const estimatedMinutes = ref(0)
const isUrgent = ref(false)
const isImportant = ref(false)
const loading = ref(false)

const submit = async () => {
  if (!title.value.trim()) return
  loading.value = true

  try {
    const task = await tasksStore.create({
      title: title.value.trim(),
      description: description.value,
      project_id: projectId.value || undefined,
      priority: priority.value,
      urgency: isUrgent.value ? 1 : 0,
      importance: isImportant.value ? 1 : 0,
      due_date: dueDate.value || undefined,
      estimated_time: estimatedMinutes.value ? estimatedMinutes.value * 60 : 0
    } as any)

    emit('created', task)
    emit('close')
  } catch {
    // error handled by store
  } finally {
    loading.value = false
  }
}
</script>

<style scoped lang="scss">
.form-group {
  margin-bottom: var(--spacing-md);

  label {
    display: block;
    font-size: var(--font-size-sm);
    font-weight: 500;
    color: var(--text-secondary);
    margin-bottom: var(--spacing-xs);
  }

  input, select, textarea {
    width: 100%;
  }
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-md);
}

.toggle-label {
  display: flex !important;
  align-items: center;
  gap: var(--spacing-sm);
  cursor: pointer;
  font-size: var(--font-size-sm);

  input[type="checkbox"] {
    width: auto;
  }
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
  margin-top: var(--spacing-lg);
  padding-top: var(--spacing-md);
  border-top: 1px solid var(--border-color);
}
</style>
