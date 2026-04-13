<template>
  <div class="github-integration">
    <div class="page-header">
      <h1>GitHub Integration</h1>
      <p class="text-secondary">Connect your GitHub account to sync issues and pull requests as tasks.</p>
    </div>

    <!-- Connection Status -->
    <div class="card connection-card">
      <div class="card-header">
        <h2>Connection</h2>
        <span
          class="connection-status"
          :class="integration.connected ? 'connected' : 'disconnected'"
        >
          {{ integration.connected ? 'Connected' : 'Not Connected' }}
        </span>
      </div>

      <div v-if="integration.connected" class="connected-info">
        <div class="user-card">
          <img
            v-if="integration.avatar_url"
            :src="integration.avatar_url"
            class="github-avatar"
            alt="GitHub Avatar"
          />
          <div class="github-user-info">
            <div class="github-username">{{ integration.username }}</div>
            <div class="github-sync-time text-xs text-tertiary">
              Last synced: {{ integration.last_synced_at ? formatDate(integration.last_synced_at) : 'Never' }}
            </div>
          </div>
          <button class="btn btn-danger btn-sm" @click="disconnect">Disconnect</button>
        </div>
      </div>

      <div v-else class="connect-form">
        <p class="text-sm text-secondary">
          Enter a GitHub Personal Access Token to connect your account.
          The token needs <code>repo</code> scope to sync issues and PRs.
        </p>
        <div class="form-group">
          <label>Personal Access Token</label>
          <input
            v-model="accessToken"
            type="password"
            placeholder="ghp_xxxxxxxxxxxxxxxxxxxx"
          />
        </div>
        <div class="form-group">
          <label>GitHub Username</label>
          <input
            v-model="username"
            type="text"
            placeholder="your-username"
          />
        </div>
        <button class="btn btn-primary" @click="connect" :disabled="!accessToken || loading">
          {{ loading ? 'Connecting...' : 'Connect GitHub' }}
        </button>
      </div>
    </div>

    <!-- Repository Tracking -->
    <div class="card" v-if="integration.connected">
      <div class="card-header">
        <h2>Tracked Repositories</h2>
        <button class="btn btn-ghost btn-sm" @click="showAddRepo = true">+ Add Repo</button>
      </div>

      <div class="repo-list">
        <div v-for="(repo, index) in repos" :key="repo.full_name" class="repo-item">
          <div class="repo-info">
            <div class="repo-name">{{ repo.full_name }}</div>
            <div class="repo-options">
              <label class="option-toggle">
                <input type="checkbox" v-model="repo.sync_issues" @change="saveRepos" />
                <span>Issues</span>
              </label>
              <label class="option-toggle">
                <input type="checkbox" v-model="repo.sync_prs" @change="saveRepos" />
                <span>PRs</span>
              </label>
            </div>
          </div>
          <div class="repo-project">
            <select v-model.number="repo.project_id" @change="saveRepos" class="project-select">
              <option :value="0">No Project</option>
              <option
                v-for="project in projectsStore.activeProjects"
                :key="project.id"
                :value="project.id"
              >
                {{ project.title }}
              </option>
            </select>
            <button class="btn btn-ghost btn-sm" @click="removeRepo(index)">Remove</button>
          </div>
        </div>

        <div v-if="repos.length === 0" class="empty-state">
          No repositories tracked. Add a repository to start syncing.
        </div>
      </div>
    </div>

    <!-- Add Repo Dialog -->
    <Teleport to="body" v-if="showAddRepo">
      <div class="modal-overlay" @click.self="showAddRepo = false">
        <div class="modal">
          <h2>Add Repository</h2>
          <div class="form-group">
            <label>Repository (owner/name)</label>
            <input
              v-model="newRepoName"
              type="text"
              placeholder="e.g. octocat/hello-world"
            />
          </div>
          <div class="form-actions">
            <button class="btn btn-ghost" @click="showAddRepo = false">Cancel</button>
            <button class="btn btn-primary" @click="addRepo" :disabled="!newRepoName">Add</button>
          </div>
        </div>
      </div>
    </Teleport>

    <!-- Synced Issues -->
    <div class="card" v-if="integration.connected && issueSyncs.length > 0">
      <div class="card-header">
        <h2>Synced Issues</h2>
      </div>
      <div class="issue-list">
        <div v-for="sync in issueSyncs" :key="sync.id" class="issue-item">
          <span class="issue-state" :class="sync.issue_state">
            {{ sync.issue_state === 'open' ? 'O' : 'C' }}
          </span>
          <div class="issue-info">
            <div class="issue-title">
              {{ sync.issue_title }}
              <span class="issue-number">#{{ sync.issue_number }}</span>
            </div>
            <div class="issue-repo text-xs text-tertiary">
              {{ sync.repo_full_name }}
              {{ sync.is_pull_request ? '(PR)' : '' }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { githubAPI } from '@/api/github'
import { useProjectsStore } from '@/stores/projects'
import type { GithubIntegration, RepoConfig, GithubIssueSync } from '@/types/models'

const projectsStore = useProjectsStore()

const integration = ref<GithubIntegration>({ connected: false })
const repos = ref<RepoConfig[]>([])
const issueSyncs = ref<GithubIssueSync[]>([])
const loading = ref(false)
const accessToken = ref('')
const username = ref('')
const showAddRepo = ref(false)
const newRepoName = ref('')

onMounted(async () => {
  await projectsStore.fetchAll()
  await fetchIntegration()
})

const fetchIntegration = async () => {
  try {
    const resp = await githubAPI.getIntegration()
    integration.value = resp.data
    repos.value = resp.data.repos || []

    if (integration.value.connected) {
      const issuesResp = await githubAPI.getIssueSyncs()
      issueSyncs.value = issuesResp.data
    }
  } catch {
    // Not connected
  }
}

const connect = async () => {
  loading.value = true
  try {
    await githubAPI.saveIntegration({
      access_token: accessToken.value,
      username: username.value
    })
    await fetchIntegration()
    accessToken.value = ''
  } catch {
    // Error handled
  } finally {
    loading.value = false
  }
}

const disconnect = async () => {
  try {
    await githubAPI.deleteIntegration()
    integration.value = { connected: false }
    repos.value = []
    issueSyncs.value = []
  } catch {
    // Error handled
  }
}

const addRepo = () => {
  const parts = newRepoName.value.split('/')
  if (parts.length !== 2) return

  repos.value.push({
    owner: parts[0],
    name: parts[1],
    full_name: newRepoName.value,
    project_id: 0,
    sync_issues: true,
    sync_prs: false
  })
  saveRepos()
  showAddRepo.value = false
  newRepoName.value = ''
}

const removeRepo = (index: number) => {
  repos.value.splice(index, 1)
  saveRepos()
}

const saveRepos = async () => {
  try {
    await githubAPI.updateRepos(repos.value)
  } catch {
    // Error handled
  }
}

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleDateString('en-US', {
    month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit'
  })
}
</script>

<style scoped lang="scss">
.github-integration {
  max-width: 800px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: var(--spacing-xl);

  h1 { margin-bottom: var(--spacing-xs); }
  p { margin: 0; }
}

.card {
  background-color: var(--bg-primary);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-sm);
  margin-bottom: var(--spacing-lg);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);

  h2 { margin: 0; font-size: var(--font-size-lg); }
}

.connection-status {
  padding: var(--spacing-xs) var(--spacing-md);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;

  &.connected { background: rgba(16, 185, 129, 0.1); color: var(--color-success); }
  &.disconnected { background: rgba(156, 163, 175, 0.1); color: var(--text-tertiary); }
}

.user-card {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.github-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
}

.github-user-info { flex: 1; }

.github-username {
  font-weight: 600;
  font-size: var(--font-size-lg);
}

.connect-form {
  p { margin-bottom: var(--spacing-lg); }
}

.form-group {
  margin-bottom: var(--spacing-md);

  label {
    display: block;
    font-size: var(--font-size-sm);
    font-weight: 500;
    color: var(--text-secondary);
    margin-bottom: var(--spacing-xs);
  }

  input, select {
    width: 100%;
  }
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
  margin-top: var(--spacing-lg);
}

.repo-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
}

.repo-item {
  padding: var(--spacing-md);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
}

.repo-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-sm);
}

.repo-name {
  font-weight: 600;
}

.repo-options {
  display: flex;
  gap: var(--spacing-md);
}

.option-toggle {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--font-size-sm);
  cursor: pointer;
}

.repo-project {
  display: flex;
  gap: var(--spacing-sm);
  align-items: center;
}

.project-select {
  flex: 1;
  font-size: var(--font-size-sm);
}

.issue-list {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-sm);
}

.issue-item {
  display: flex;
  gap: var(--spacing-md);
  padding: var(--spacing-sm) 0;
  border-bottom: 1px solid var(--border-light);
}

.issue-state {
  width: 20px;
  height: 20px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  font-weight: 700;
  flex-shrink: 0;

  &.open { background: rgba(16, 185, 129, 0.15); color: var(--color-success); }
  &.closed { background: rgba(139, 92, 246, 0.15); color: var(--color-secondary); }
}

.issue-title {
  font-size: var(--font-size-sm);
  font-weight: 500;
}

.issue-number {
  color: var(--text-tertiary);
  font-weight: 400;
}

.empty-state {
  text-align: center;
  padding: var(--spacing-lg);
  color: var(--text-tertiary);
  font-size: var(--font-size-sm);
}
</style>
