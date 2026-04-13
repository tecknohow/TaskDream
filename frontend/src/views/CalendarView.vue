<template>
  <div class="calendar-view">
    <div class="calendar-header">
      <button @click="previousMonth" class="btn btn-ghost">← Previous</button>
      <h1>{{ monthYear }}</h1>
      <button @click="nextMonth" class="btn btn-ghost">Next →</button>
    </div>

    <div class="calendar-grid">
      <div v-for="day in ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']" :key="day" class="weekday-header">
        {{ day }}
      </div>

      <div
        v-for="day in calendarDays"
        :key="`${day.year}-${day.month}-${day.date}`"
        class="calendar-day"
        :class="{ 'other-month': day.month !== currentMonth, today: isToday(day) }"
      >
        <div class="day-number">{{ day.date }}</div>
        <div class="day-tasks">
          <div
            v-for="task in getTasksForDay(day)"
            :key="task.id"
            class="day-task"
            @click="selectTask(task)"
          >
            <span class="task-title">{{ task.title }}</span>
          </div>
          <div v-if="getTasksForDay(day).length === 0" class="no-tasks"></div>
        </div>
      </div>
    </div>

    <div v-if="tasksStore.selectedTask" class="task-detail-sidebar">
      <h2>{{ tasksStore.selectedTask.title }}</h2>
      <p v-if="tasksStore.selectedTask.description" class="task-description">
        {{ tasksStore.selectedTask.description }}
      </p>
      <div class="task-meta">
        <div>Status: <strong>{{ tasksStore.selectedTask.status }}</strong></div>
        <div>Priority: <strong>{{ priorityLabel(tasksStore.selectedTask.priority) }}</strong></div>
        <div v-if="tasksStore.selectedTask.dueDate">
          Due: <strong>{{ formatDate(tasksStore.selectedTask.dueDate) }}</strong>
        </div>
      </div>
      <button class="btn btn-primary btn-sm" @click="navigateToTask">View Details</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTasksStore } from '@/stores/tasks'
import { getDaysInMonth, getFirstDayOfMonth } from '@/utils/date'
import { formatDate } from '@/utils/format'
import type { Task } from '@/types/models'

const router = useRouter()
const tasksStore = useTasksStore()

const currentDate = ref(new Date())

const currentMonth = computed(() => currentDate.value.getMonth())
const currentYear = computed(() => currentDate.value.getFullYear())

const monthYear = computed(() => {
  const months = ['January', 'February', 'March', 'April', 'May', 'June',
    'July', 'August', 'September', 'October', 'November', 'December']
  return `${months[currentMonth.value]} ${currentYear.value}`
})

const calendarDays = computed(() => {
  const days = []
  const firstDay = getFirstDayOfMonth(currentDate.value)
  const daysInMonth = getDaysInMonth(currentDate.value)
  const previousMonthDays = getDaysInMonth(new Date(currentYear.value, currentMonth.value - 1))

  // Previous month days
  for (let i = firstDay - 1; i >= 0; i--) {
    days.push({
      date: previousMonthDays - i,
      month: currentMonth.value - 1,
      year: currentMonth.value === 0 ? currentYear.value - 1 : currentYear.value
    })
  }

  // Current month days
  for (let i = 1; i <= daysInMonth; i++) {
    days.push({
      date: i,
      month: currentMonth.value,
      year: currentYear.value
    })
  }

  // Next month days
  const remaining = 42 - days.length
  for (let i = 1; i <= remaining; i++) {
    days.push({
      date: i,
      month: currentMonth.value + 1,
      year: currentMonth.value === 11 ? currentYear.value + 1 : currentYear.value
    })
  }

  return days
})

onMounted(() => {
  tasksStore.fetchAll()
})

const previousMonth = () => {
  currentDate.value = new Date(currentDate.value.getFullYear(), currentDate.value.getMonth() - 1)
}

const nextMonth = () => {
  currentDate.value = new Date(currentDate.value.getFullYear(), currentDate.value.getMonth() + 1)
}

const isToday = (day: any) => {
  const today = new Date()
  return day.date === today.getDate() &&
    day.month === today.getMonth() &&
    day.year === today.getFullYear()
}

const getTasksForDay = (day: any) => {
  return tasksStore.tasks.filter(task => {
    if (!task.dueDate) return false
    const dueDate = new Date(task.dueDate)
    return dueDate.getDate() === day.date &&
      dueDate.getMonth() === day.month &&
      dueDate.getFullYear() === day.year
  }).slice(0, 2)
}

const selectTask = (task: Task) => {
  tasksStore.selectTask(task)
}

const navigateToTask = () => {
  if (tasksStore.selectedTask) {
    router.push(`/tasks/${tasksStore.selectedTask.id}`)
  }
}

const priorityLabel = (p: number) => {
  const labels = ['Low', 'Medium', 'High', 'Urgent']
  return labels[p] || 'Unknown'
}
</script>

<style scoped lang="scss">
.calendar-view {
  max-width: 1200px;
  margin: 0 auto;
}

.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-2xl);

  h1 {
    margin: 0;
    min-width: 250px;
    text-align: center;
  }
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: var(--spacing-sm);
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-md);
  margin-bottom: var(--spacing-xl);
}

.weekday-header {
  padding: var(--spacing-md);
  text-align: center;
  font-weight: 600;
  color: var(--text-secondary);
  border-bottom: 2px solid var(--border-color);
}

.calendar-day {
  min-height: 100px;
  padding: var(--spacing-sm);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  background-color: var(--bg-secondary);
  display: flex;
  flex-direction: column;

  &.today {
    background-color: rgba(37, 99, 235, 0.05);
    border-color: var(--color-primary);
  }

  &.other-month {
    opacity: 0.5;
  }
}

.day-number {
  font-weight: 600;
  font-size: var(--font-size-sm);
  margin-bottom: var(--spacing-xs);
}

.day-tasks {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.day-task {
  background-color: var(--color-primary);
  color: white;
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  cursor: pointer;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  transition: transform var(--transition-fast);

  &:hover {
    transform: translateY(-2px);
  }
}

.no-tasks {
  flex: 1;
}

.task-detail-sidebar {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  max-width: 400px;
  margin-left: auto;

  h2 {
    margin-top: 0;
  }
}

.task-description {
  color: var(--text-secondary);
  margin-bottom: var(--spacing-md);
  line-height: 1.5;
}

.task-meta {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-lg);
  font-size: var(--font-size-sm);

  div {
    padding: var(--spacing-sm);
    background-color: var(--bg-secondary);
    border-radius: var(--radius-md);
  }
}

@media (max-width: 768px) {
  .calendar-grid {
    gap: 2px;
    padding: 2px;
  }

  .calendar-day {
    min-height: 80px;
    padding: 4px;
  }

  .day-number {
    font-size: var(--font-size-xs);
  }

  .task-detail-sidebar {
    max-width: none;
    margin-left: 0;
    margin-top: var(--spacing-lg);
  }
}
</style>
