import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { projectsAPI } from '@/api/projects'
import type { Project, Bucket } from '@/types/models'

export const useProjectsStore = defineStore('projects', () => {
  const projects = ref<Project[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const selectedProject = ref<Project | null>(null)

  const activeProjects = computed(() => {
    return projects.value.filter(p => !p.is_archived)
  })

  const fetchAll = async () => {
    loading.value = true
    error.value = null

    try {
      const response = await projectsAPI.getAll()
      projects.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch projects'
    } finally {
      loading.value = false
    }
  }

  const fetchById = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      const response = await projectsAPI.getById(id)
      const index = projects.value.findIndex(p => p.id === id)
      if (index > -1) {
        projects.value[index] = response.data
      } else {
        projects.value.push(response.data)
      }
      selectedProject.value = response.data
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch project'
    } finally {
      loading.value = false
    }
  }

  const create = async (project: Partial<Project>) => {
    loading.value = true
    error.value = null

    try {
      const response = await projectsAPI.create(project)
      projects.value.push(response.data)
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to create project'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const update = async (id: number, project: Partial<Project>) => {
    loading.value = true
    error.value = null

    try {
      const response = await projectsAPI.update(id, project)
      const index = projects.value.findIndex(p => p.id === id)
      if (index > -1) {
        projects.value[index] = response.data
      }
      if (selectedProject.value?.id === id) {
        selectedProject.value = response.data
      }
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to update project'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const deleteProject = async (id: number) => {
    loading.value = true
    error.value = null

    try {
      await projectsAPI.delete(id)
      projects.value = projects.value.filter(p => p.id !== id)
      if (selectedProject.value?.id === id) {
        selectedProject.value = null
      }
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to delete project'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const selectProject = (project: Project | null) => {
    selectedProject.value = project
  }

  return {
    projects,
    selectedProject,
    loading,
    error,
    activeProjects,
    fetchAll,
    fetchById,
    create,
    update,
    deleteProject,
    selectProject
  }
})
