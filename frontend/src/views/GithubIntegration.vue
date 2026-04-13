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
						v-if="integration.avatarUrl"
						:src="integration.avatarUrl"
						class="github-avatar"
						alt="GitHub Avatar"
					/>
					<div class="github-user-info">
						<div class="github-username">{{ integration.username }}</div>
						<div class="github-sync-time text-xs text-tertiary">
							Last synced: {{ integration.lastSyncedAt ? formatDate(integration.lastSyncedAt) : 'Never' }}
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
		<div v-if="integration.connected" class="card">
			<div class="card-header">
				<h2>Tracked Repositories</h2>
				<button class="btn btn-ghost btn-sm" @click="showAddRepo = true">+ Add Repo</button>
			</div>

			<div class="repo-list">
				<div v-for="(repo, index) in repos" :key="repo.fullName" class="repo-item">
					<div class="repo-info">
						<div class="repo-name">{{ repo.fullName }}</div>
						<div class="repo-options">
							<label class="option-toggle">
								<input type="checkbox" v-model="repo.syncIssues" @change="saveRepos" />
								<span>Issues</span>
							</label>
							<label class="option-toggle">
								<input type="checkbox" v-model="repo.syncPrs" @change="saveRepos" />
								<span>PRs</span>
							</label>
						</div>
					</div>
					<div class="repo-project">
						<button class="btn btn-ghost btn-sm" @click="removeRepo(index)">Remove</button>
					</div>
				</div>

				<div v-if="repos.length === 0" class="empty-state">
					No repositories tracked. Add a repository to start syncing.
				</div>
			</div>
		</div>

		<!-- Add Repo Dialog -->
		<Teleport to="body">
			<div v-if="showAddRepo" class="modal-overlay" @click.self="showAddRepo = false">
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
		<div v-if="integration.connected && issueSyncs.length > 0" class="card">
			<div class="card-header">
				<h2>Synced Issues</h2>
			</div>
			<div class="issue-list">
				<div v-for="sync in issueSyncs" :key="sync.id" class="issue-item">
					<span class="issue-state" :class="sync.issueState">
						{{ sync.issueState === 'open' ? 'O' : 'C' }}
					</span>
					<div class="issue-info">
						<div class="issue-title">
							{{ sync.issueTitle }}
							<span class="issue-number">#{{ sync.issueNumber }}</span>
						</div>
						<div class="issue-repo text-xs text-tertiary">
							{{ sync.repoFullName }}
							{{ sync.isPullRequest ? '(PR)' : '' }}
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts" setup>
import {ref, onMounted} from 'vue'

import {AuthenticatedHTTPFactory} from '@/helpers/fetcher'

interface GithubIntegrationData {
	connected: boolean
	username?: string
	avatarUrl?: string
	lastSyncedAt?: string
}

interface RepoConfig {
	owner: string
	name: string
	fullName: string
	projectId: number
	syncIssues: boolean
	syncPrs: boolean
}

interface GithubIssueSync {
	id: number
	issueTitle: string
	issueNumber: number
	issueState: string
	repoFullName: string
	isPullRequest: boolean
}

const http = AuthenticatedHTTPFactory()

const integration = ref<GithubIntegrationData>({connected: false})
const repos = ref<RepoConfig[]>([])
const issueSyncs = ref<GithubIssueSync[]>([])
const loading = ref(false)
const accessToken = ref('')
const username = ref('')
const showAddRepo = ref(false)
const newRepoName = ref('')

onMounted(async () => {
	await fetchIntegration()
})

async function fetchIntegration() {
	try {
		const resp = await http.get('/integrations/github')
		integration.value = resp.data
		repos.value = resp.data.repos || []

		if (integration.value.connected) {
			const issuesResp = await http.get('/integrations/github/issues')
			issueSyncs.value = issuesResp.data || []
		}
	} catch {
		// Not connected
	}
}

async function connect() {
	loading.value = true
	try {
		await http.put('/integrations/github', {
			access_token: accessToken.value,
			username: username.value,
		})
		await fetchIntegration()
		accessToken.value = ''
	} catch {
		// Error handled
	} finally {
		loading.value = false
	}
}

async function disconnect() {
	try {
		await http.delete('/integrations/github')
		integration.value = {connected: false}
		repos.value = []
		issueSyncs.value = []
	} catch {
		// Error handled
	}
}

function addRepo() {
	const parts = newRepoName.value.split('/')
	if (parts.length !== 2) return

	repos.value.push({
		owner: parts[0],
		name: parts[1],
		fullName: newRepoName.value,
		projectId: 0,
		syncIssues: true,
		syncPrs: false,
	})
	saveRepos()
	showAddRepo.value = false
	newRepoName.value = ''
}

function removeRepo(index: number) {
	repos.value.splice(index, 1)
	saveRepos()
}

async function saveRepos() {
	try {
		await http.post('/integrations/github/repos', {repos: repos.value})
	} catch {
		// Error handled
	}
}

function formatDate(dateStr: string) {
	return new Date(dateStr).toLocaleDateString('en-US', {
		month: 'short',
		day: 'numeric',
		hour: '2-digit',
		minute: '2-digit',
	})
}
</script>

<style scoped lang="scss">
.github-integration {
	max-width: 800px;
	margin: 0 auto;
}

.page-header {
	margin-bottom: 2rem;

	h1 { margin-bottom: .25rem; }
	p { margin: 0; }
}

.card {
	background-color: var(--white);
	border: 1px solid var(--grey-200);
	border-radius: $radius;
	padding: 1.5rem;
	box-shadow: var(--shadow-sm);
	margin-bottom: 1.5rem;
}

.card-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: 1.5rem;

	h2 { margin: 0; font-size: 1.125rem; }
}

.connection-status {
	padding: .25rem .75rem;
	border-radius: $radius;
	font-size: .75rem;
	font-weight: 600;

	&.connected { background: rgba(16, 185, 129, 0.1); color: var(--success); }
	&.disconnected { background: rgba(156, 163, 175, 0.1); color: var(--grey-500); }
}

.user-card {
	display: flex;
	align-items: center;
	gap: .75rem;
}

.github-avatar {
	width: 48px;
	height: 48px;
	border-radius: 50%;
}

.github-user-info { flex: 1; }

.github-username {
	font-weight: 600;
	font-size: 1.125rem;
}

.connect-form {
	p { margin-bottom: 1.5rem; }
}

.form-group {
	margin-bottom: .75rem;

	label {
		display: block;
		font-size: .875rem;
		font-weight: 500;
		color: var(--grey-500);
		margin-bottom: .25rem;
	}

	input, select {
		width: 100%;
	}
}

.form-actions {
	display: flex;
	justify-content: flex-end;
	gap: .75rem;
	margin-top: 1rem;
}

.repo-list {
	display: flex;
	flex-direction: column;
	gap: .75rem;
}

.repo-item {
	padding: .75rem;
	border: 1px solid var(--grey-200);
	border-radius: $radius;
}

.repo-info {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: .5rem;
}

.repo-name {
	font-weight: 600;
}

.repo-options {
	display: flex;
	gap: .75rem;
}

.option-toggle {
	display: flex;
	align-items: center;
	gap: .25rem;
	font-size: .875rem;
	cursor: pointer;
}

.repo-project {
	display: flex;
	gap: .5rem;
	align-items: center;
	justify-content: flex-end;
}

.issue-list {
	display: flex;
	flex-direction: column;
	gap: .5rem;
}

.issue-item {
	display: flex;
	gap: .75rem;
	padding: .5rem 0;
	border-bottom: 1px solid var(--grey-100);
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

	&.open { background: rgba(16, 185, 129, 0.15); color: var(--success); }
	&.closed { background: rgba(139, 92, 246, 0.15); color: var(--primary); }
}

.issue-title {
	font-size: .875rem;
	font-weight: 500;
}

.issue-number {
	color: var(--grey-400);
	font-weight: 400;
}

.empty-state {
	text-align: center;
	padding: 1.5rem;
	color: var(--grey-400);
	font-size: .875rem;
}

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
	z-index: 100;
}

.modal {
	background-color: var(--white);
	border-radius: $radius;
	padding: 1.5rem;
	width: 400px;
	max-width: 90vw;
	box-shadow: var(--shadow-lg);

	h2 {
		margin-bottom: 1rem;
	}
}
</style>
