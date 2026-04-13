<template>
  <div class="notes-view">
    <div class="notes-header">
      <h1>Notes</h1>
      <button class="btn btn-primary" @click="createNewNote">+ New Note</button>
    </div>

    <div class="notes-container">
      <div class="notes-list">
        <div
          v-for="note in notes"
          :key="note.id"
          class="note-item"
          :class="{ active: selectedNote?.id === note.id }"
          @click="selectedNote = note"
        >
          <h3>{{ note.title }}</h3>
          <p class="note-preview">
            {{ note.content.substring(0, 100) }}...
          </p>
          <div class="note-date">
            {{ formatDate(note.updatedAt) }}
          </div>
        </div>

        <div v-if="notes.length === 0" class="empty-notes">
          No notes yet. Create one to get started!
        </div>
      </div>

      <div v-if="selectedNote" class="note-editor">
        <div class="editor-header">
          <input
            v-model="selectedNote.title"
            type="text"
            class="note-title-input"
            placeholder="Note title"
            @change="saveNote"
          />
          <button class="btn btn-danger btn-sm" @click="deleteNote">Delete</button>
        </div>

        <textarea
          v-model="selectedNote.content"
          class="note-content-input"
          placeholder="Start typing..."
          @change="saveNote"
        ></textarea>
      </div>

      <div v-else class="no-selection">
        Select a note to edit
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { formatDate } from '@/utils/format'
import type { Note } from '@/types/models'

const notes = ref<Note[]>([
  {
    id: '1',
    title: 'Sample Note',
    content: 'This is a sample note. You can create notes to keep track of ideas and information related to your projects.',
    createdAt: new Date(),
    updatedAt: new Date(),
    userId: 'user1'
  }
])

const selectedNote = ref<Note | null>(null)

onMounted(() => {
  if (notes.value.length > 0) {
    selectedNote.value = notes.value[0]
  }
})

const createNewNote = () => {
  const newNote: Note = {
    id: String(Date.now()),
    title: 'Untitled Note',
    content: '',
    createdAt: new Date(),
    updatedAt: new Date(),
    userId: 'user1'
  }
  notes.value.unshift(newNote)
  selectedNote.value = newNote
}

const saveNote = async () => {
  if (selectedNote.value) {
    selectedNote.value.updatedAt = new Date()
    const index = notes.value.findIndex(n => n.id === selectedNote.value!.id)
    if (index > -1) {
      notes.value[index] = selectedNote.value
    }
  }
}

const deleteNote = async () => {
  if (selectedNote.value && confirm('Delete this note?')) {
    notes.value = notes.value.filter(n => n.id !== selectedNote.value!.id)
    selectedNote.value = notes.value.length > 0 ? notes.value[0] : null
  }
}
</script>

<style scoped lang="scss">
.notes-view {
  max-width: 1200px;
  margin: 0 auto;
}

.notes-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-2xl);

  h1 {
    margin: 0;
  }
}

.notes-container {
  display: grid;
  grid-template-columns: 300px 1fr;
  gap: var(--spacing-lg);
  min-height: 600px;
}

.notes-list {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-md);
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.note-item {
  padding: var(--spacing-md);
  border-radius: var(--radius-md);
  background-color: var(--bg-secondary);
  cursor: pointer;
  transition: all var(--transition-fast);
  border: 2px solid transparent;

  h3 {
    margin: 0 0 var(--spacing-xs) 0;
    font-size: var(--font-size-base);
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  &:hover {
    background-color: var(--bg-tertiary);
  }

  &.active {
    background-color: var(--color-primary);
    color: white;
    border-color: var(--color-primary);

    .note-preview,
    .note-date {
      color: rgba(255, 255, 255, 0.8);
    }
  }
}

.note-preview {
  font-size: var(--font-size-sm);
  color: var(--text-secondary);
  margin: 0 0 var(--spacing-xs) 0;
  line-height: 1.3;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.note-date {
  font-size: var(--font-size-xs);
  color: var(--text-tertiary);
}

.empty-notes {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: var(--text-tertiary);
  text-align: center;
  padding: var(--spacing-lg);
}

.note-editor {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.editor-header {
  display: flex;
  gap: var(--spacing-md);
  align-items: flex-start;
}

.note-title-input {
  flex: 1;
  font-size: var(--font-size-lg);
  font-weight: 600;
  border: 1px solid var(--border-color);
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-md);

  &:focus {
    outline: none;
    border-color: var(--color-primary);
  }
}

.note-content-input {
  flex: 1;
  padding: var(--spacing-md);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  font-family: inherit;
  font-size: var(--font-size-base);
  color: var(--text-primary);
  background-color: var(--bg-secondary);
  resize: vertical;
  min-height: 400px;

  &:focus {
    outline: none;
    border-color: var(--color-primary);
    background-color: var(--bg-primary);
  }
}

.no-selection {
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  color: var(--text-tertiary);
  font-size: var(--font-size-lg);
}

@media (max-width: 768px) {
  .notes-container {
    grid-template-columns: 1fr;
  }

  .notes-list {
    max-height: 300px;
  }

  .note-content-input {
    min-height: 200px;
  }
}
</style>
