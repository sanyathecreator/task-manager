document.addEventListener('DOMContentLoaded', () => {
    const API = 'https://localhost:8080'
    const taskInput = document.getElementById('task-input');
    const addTaskBtn = document.getElementById('add-task-btn');
    const taskList = document.getElementById('task-list');

    function renderTask(task) {
        const li = document.createElement('li');
        li.dataset.id = task.id;

        const checkbox = document.createElement('input');
        checkbox.type = 'checkbox';
        checkbox.checked = task.completed;
        checkbox.addEventListener('change', async () => {
            await fetch(`${API}/tasks/${task.id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ title: task.title, completed: checkbox.checked })
            });
        });

        const deleteBtn = document.createElement('btn');
        deleteBtn.textContent = 'Delete';
        deleteBtn.addEventListener('click', async () => {
            await fetch(`${API}/tasks/${task.id}`, { method: 'DELETE' });
            li.remove();
        })

        li.appendChild(checkbox);
        li.appendChild(document.createTextNode(task.title));
        li.appendChild(deleteBtn);
        taskList.appendChild(li);
    }

    async function loadTasks() {
        const res = await fetch(`${API}/tasks`);
        const tasks = await res.json();
        tasks.forEach(renderTask);
    }

    const addTask = async (event) => {
        event.preventDefault();
        const taskTitle = taskInput.value.trim();
        if (!taskTitle) return;

        const res = await fetch(`${API}/tasks`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ title: taskTitle, completed: false })
        })

        const data = await res.json();
        renderTask(data.task);
        taskInput.value = '';
    }

    addTaskBtn.addEventListener('click', addTask);
    taskInput.addEventListener('keypress', (e) => {
        if (e.key === 'Enter') addTask(e);
    })

    loadTasks();
}); 
