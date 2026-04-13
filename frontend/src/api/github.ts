import client from './client'
import type { GithubIntegration, GithubIssueSync, RepoConfig } from '@/types/models'

export const githubAPI = {
  getIntegration: () => {
    return client.get<GithubIntegration>('/integrations/github')
  },

  saveIntegration: (data: { access_token: string; username?: string; avatar_url?: string }) => {
    return client.post('/integrations/github', data)
  },

  deleteIntegration: () => {
    return client.delete('/integrations/github')
  },

  updateRepos: (repos: RepoConfig[]) => {
    return client.put('/integrations/github/repos', { repos })
  },

  getIssueSyncs: (params?: { task_id?: number; repo?: string }) => {
    return client.get<GithubIssueSync[]>('/integrations/github/issues', { params })
  }
}
