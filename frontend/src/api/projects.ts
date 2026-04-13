import client from './client'
import type { Project, Bucket } from '@/types/models'

export const projectsAPI = {
  getAll: () => {
    return client.get<Project[]>('/projects')
  },

  getById: (id: string) => {
    return client.get<Project>(`/projects/${id}`)
  },

  create: (project: Partial<Project>) => {
    return client.post<Project>('/projects', project)
  },

  update: (id: string, project: Partial<Project>) => {
    return client.put<Project>(`/projects/${id}`, project)
  },

  delete: (id: string) => {
    return client.delete(`/projects/${id}`)
  },

  addMember: (projectId: string, userId: string) => {
    return client.post(`/projects/${projectId}/members`, { userId })
  },

  removeMember: (projectId: string, userId: string) => {
    return client.delete(`/projects/${projectId}/members/${userId}`)
  },

  getBuckets: (projectId: string) => {
    return client.get<Bucket[]>(`/projects/${projectId}/buckets`)
  },

  createBucket: (projectId: string, bucket: Partial<Bucket>) => {
    return client.post<Bucket>(`/projects/${projectId}/buckets`, bucket)
  },

  updateBucket: (projectId: string, bucketId: string, bucket: Partial<Bucket>) => {
    return client.put<Bucket>(`/projects/${projectId}/buckets/${bucketId}`, bucket)
  },

  deleteBucket: (projectId: string, bucketId: string) => {
    return client.delete(`/projects/${projectId}/buckets/${bucketId}`)
  }
}
