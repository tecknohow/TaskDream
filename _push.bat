@echo off
cd /d D:\devprojs\GitHub\TaskDream
git add -A
git commit -m "feat: initial TaskDream scaffold - Go backend + Vue 3 frontend"
git push -u origin main
echo PUSH_COMPLETE > D:\devprojs\GitHub\TaskDream\.push_status
