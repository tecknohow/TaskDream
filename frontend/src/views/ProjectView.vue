<template>
  <div class="project-view" v-if="project">
    <div class="project-header">
      <div class="project-title">
        <div class="project-color" :style="{ backgroundColor: project.color || 'var(--color-primary)' }"></div>
        <h1>{{ project.name }}</h1>
      </div>
      <button class="btn btn-primary" @click="showCreateTask = true">+ Add Task</button>
    </div>

    <div class="project-description" v-if="project.description">
      {{ project.description }}
    </div>

    <div class="tasks-section">
      <h2>Tasks</h2>
      <div class="tasks-grid">
        <div
          v-for="bucket in project.buckets"
          :key="bucket.id"
          class="bucket"
        >
          <h3>{{ bucket.title }}</h3>
          <div class="task-list">
            <TaskItem
              v-for="task in bucket.tasks"
              :key="task.id"
              :task="task"
              @click="selectTask(task)"
            />
            <div v-if="bucket.tasks.length === 0" class="empty-bucket">
              No tasks in this bucket
            </div>
          </div>
        </div>
      </div>
    </div>

    <CreateTaskDialog
      v-if="showCreateTask"
      :project-id="project.id"
      @close="showCreateTask = false"
    />
  </div>

  <div v-else class="loading">
    Loading project...
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useProjectsStore } from '@/stores/projects'
import TaskItem from '@/components/TaskItem.vue'
import CreateTaskDialog from '@/components/CreateTaskDialog.vue'
import { useTasksStore } from '@/stores/tasks'
import type { Task } from '@/types/models'

const route = useRoute()
const projectsStore = useProjectsStore()
const tasksStore = useTasksStore()
const showCreateTask = ref(false)

const project = ref(projectsStore.selectedProject)

onMounted(async () => {
  const projectId = route.params.id as string
  await projectsStore.fetchById(projectId)
  await projectsStore.fetchBuckets(projectId)
  project.value = projectsStore.selectedProject
})

const selectTask = (task: Task) => {
  tasksStore.selectTask(task)
}
</script>

<style scoped lang="scss">
.project-view {
  max-width: 1200px;
  margin: 0 auto;
}

.project-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-2xl);
}

.project-title {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);

  h1 {
    margin: 0;
  }
}

.project-color {
  width: 40px;
  height: 40px;
  border-radius: var(--radius-md);
}

.project-description {
  padding: var(--spacing-md) var(--spacing-lg);
  background-color: var(--bg-tertiary);
  border-radius: var(--radius-md);
  margin-bottom: var(--spacing-xl);
  color: var(--text-secondary);
}

.tasks-section {
  h2 {
    margin-bottom: var(--spacing-lg);
  }
}

.tasks-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: var(--spacing-lg);
}

.bucket {
  background-color: var(--bg-secondary);
  border-radius: var(--radius-lg);
  padding: var(--spacing-md);

  h3 {
    margin-bottom: var(--spacing-md);
    font-size: var(--font-size-base);
  }
}

.task-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.empty-bucket {
  text-align: center;
  padding: var(--spacing-md);
  color: var(--text-tertiary);
  font-size: var(--font-size-sm);
}

.loading {
  text-align: center;
  padding: var(--spacing-xl);
  color: var(--text-secondary);
}
</style>
