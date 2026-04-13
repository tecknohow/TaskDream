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
    return projects.value.filter(p => !p.archived)
  })

  const fetchAll = async () => {
    loading.value = true
    error.value = null

    try {
      const response = await projectsAPI.getAll()
      projects.value = response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to fetch projects'
    } finally {
      loading.value = false
    }
  }

  const fetchById = async (id: string) => {
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
      error.value = err.response?.data?.message || 'Failed to fetch project'
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
      error.value = err.response?.data?.message || 'Failed to create project'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const update = async (id: string, project: Partial<Project>) => {
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
      error.value = err.response?.data?.message || 'Failed to update project'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const deleteProject = async (id: string) => {
    loading.value = true
    error.value = null

    try {
      await projectsAPI.delete(id)
      projects.value = projects.value.filter(p => p.id !== id)
      if (selectedProject.value?.id === id) {
        selectedProject.value = null
      }
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to delete project'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  const fetchBuckets = async (projectId: string) => {
    try {
      const response = await projectsAPI.getBuckets(projectId)
      const project = projects.value.find(p => p.id === projectId)
      if (project) {
        project.buckets = response.data
      }
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to fetch buckets'
    }
  }

  const createBucket = async (projectId: string, bucket: Partial<Bucket>) => {
    try {
      const response = await projectsAPI.createBucket(projectId, bucket)
      const project = projects.value.find(p => p.id === projectId)
      if (project && project.buckets) {
        project.buckets.push(response.data)
      }
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to create bucket'
      throw error.value
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
    fetchBuckets,
    createBucket,
    selectProject
  }
})
