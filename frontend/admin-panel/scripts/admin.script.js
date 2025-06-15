const API_URL = 'http://localhost:8080';

// Загрузка данных при загрузке страницы
document.addEventListener('DOMContentLoaded', () => {
    loadTeachers();
    loadSubjects();
    loadGroups();
    loadTSGs();
});

// Показать форму для добавления
function showAddForm(type) {
    document.getElementById(`${type}-form`).style.display = 'block';
    if (type === 'tsg') {
        populateSelects();
    }
}

// Скрыть форму
function hideForm(type) {
    document.getElementById(`${type}-form`).style.display = 'none';
}

// Загрузка преподавателей
async function loadTeachers() {
    try {
        const response = await fetch(`${API_URL}/teachers`);
        if (!response.ok) {
            console.error('Error fetching teachers:', response.status, response.statusText);
            alert('Ошибка загрузки преподавателей');
            return;
        }
        const teachers = await response.json();
        console.log('Teachers data:', teachers); // Отладка
        const table = document.getElementById('teachers-table');
        table.innerHTML = '';
        teachers.forEach(teacher => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${teacher.id || 'N/A'}</td>
                <td>${teacher.fullName || 'N/A'}</td>
                <td>
                    <button class="edit" onclick="editTeacher(${teacher.id || 0}, '${teacher.fullName || ''}')">Редактировать</button>
                    <button class="delete" onclick="deleteTeacher(${teacher.id || 0})">Удалить</button>
                </td>
            `;
            table.appendChild(row);
        });
    } catch (error) {
        console.error('Error in loadTeachers:', error);
        alert('Ошибка загрузки преподавателей');
    }
}

// Добавление преподавателя
async function addTeacher() {
    const name = document.getElementById('teacher-name').value;
    if (!name) {
        alert('Введите ФИО преподавателя');
        return;
    }
    try {
        await fetch(`${API_URL}/teachers`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ fullName: name })
        });
        hideForm('teacher');
        document.getElementById('teacher-name').value = '';
        loadTeachers();
    } catch (error) {
        console.error('Error adding teacher:', error);
        alert('Ошибка добавления преподавателя');
    }
}

// Редактирование преподавателя
async function editTeacher(id, currentName) {
    const name = prompt('Введите новое ФИО:', currentName);
    if (name) {
        try {
            await fetch(`${API_URL}/teachers/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ fullName: name })
            });
            loadTeachers();
        } catch (error) {
            console.error('Error editing teacher:', error);
            alert('Ошибка редактирования преподавателя');
        }
    }
}

// Удаление преподавателя
async function deleteTeacher(id) {
    if (confirm('Вы уверены, что хотите удалить преподавателя?')) {
        try {
            await fetch(`${API_URL}/teachers/${id}`, {
                method: 'DELETE'
            });
            loadTeachers();
        } catch (error) {
            console.error('Error deleting teacher:', error);
            alert('Ошибка удаления преподавателя');
        }
    }
}

// Загрузка предметов
async function loadSubjects() {
    try {
        const response = await fetch(`${API_URL}/subjects`);
        if (!response.ok) {
            console.error('Error fetching subjects:', response.status, response.statusText);
            alert('Ошибка загрузки предметов');
            return;
        }
        const subjects = await response.json();
        console.log('Subjects data:', subjects); // Отладка
        const table = document.getElementById('subjects-table');
        table.innerHTML = '';
        subjects.forEach(subject => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${subject.id || 'N/A'}</td>
                <td>${subject.name || 'N/A'}</td>
                <td>
                    <button class="edit" onclick="editSubject(${subject.id || 0}, '${subject.name || ''}')">Редактировать</button>
                    <button class="delete" onclick="deleteSubject(${subject.id || 0})">Удалить</button>
                </td>
            `;
            table.appendChild(row);
        });
    } catch (error) {
        console.error('Error in loadSubjects:', error);
        alert('Ошибка загрузки предметов');
    }
}

// Добавление предмета
async function addSubject() {
    const name = document.getElementById('subject-name').value;
    if (!name) {
        alert('Введите название предмета');
        return;
    }
    try {
        await fetch(`${API_URL}/subjects`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name })
        });
        hideForm('subject');
        document.getElementById('subject-name').value = '';
        loadSubjects();
    } catch (error) {
        console.error('Error adding subject:', error);
        alert('Ошибка добавления предмета');
    }
}

// Редактирование предмета
async function editSubject(id, currentName) {
    const name = prompt('Введите новое название предмета:', currentName);
    if (name) {
        try {
            await fetch(`${API_URL}/subjects/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name })
            });
            loadSubjects();
        } catch (error) {
            console.error('Error editing subject:', error);
            alert('Ошибка редактирования предмета');
        }
    }
}

// Удаление предмета
async function deleteSubject(id) {
    if (confirm('Вы уверены, что хотите удалить предмет?')) {
        try {
            await fetch(`${API_URL}/subjects/${id}`, {
                method: 'DELETE'
            });
            loadSubjects();
        } catch (error) {
            console.error('Error deleting subject:', error);
            alert('Ошибка удаления предмета');
        }
    }
}

// Загрузка групп
async function loadGroups() {
    try {
        const response = await fetch(`${API_URL}/groups`);
        if (!response.ok) {
            console.error('Error fetching groups:', response.status, response.statusText);
            alert('Ошибка загрузки групп');
            return;
        }
        const groups = await response.json();
        console.log('Groups data:', groups); // Отладка
        const table = document.getElementById('groups-table');
        table.innerHTML = '';
        groups.forEach(group => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${group.id || 'N/A'}</td>
                <td>${group.name || 'N/A'}</td>
                <td>
                    <button class="edit" onclick="editGroup(${group.id || 0}, '${group.name || ''}')">Редактировать</button>
                    <button class="delete" onclick="deleteGroup(${group.id || 0})">Удалить</button>
                </td>
            `;
            table.appendChild(row);
        });
    } catch (error) {
        console.error('Error in loadGroups:', error);
        alert('Ошибка загрузки групп');
    }
}

// Добавление группы
async function addGroup() {
    const name = document.getElementById('group-name').value;
    if (!name) {
        alert('Введите название группы');
        return;
    }
    try {
        await fetch(`${API_URL}/groups`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name })
        });
        hideForm('group');
        document.getElementById('group-name').value = '';
        loadGroups();
    } catch (error) {
        console.error('Error adding group:', error);
        alert('Ошибка добавления группы');
    }
}

// Редактирование группы
async function editGroup(id, currentName) {
    const name = prompt('Введите новое название группы:', currentName);
    if (name) {
        try {
            await fetch(`${API_URL}/groups/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ name })
            });
            loadGroups();
        } catch (error) {
            console.error('Error editing group:', error);
            alert('Ошибка редактирования группы');
        }
    }
}

// Удаление группы
async function deleteGroup(id) {
    if (confirm('Вы уверены, что хотите удалить группу?')) {
        try {
            await fetch(`${API_URL}/groups/${id}`, {
                method: 'DELETE'
            });
            loadGroups();
        } catch (error) {
            console.error('Error deleting group:', error);
            alert('Ошибка удаления группы');
        }
    }
}

// Загрузка связей
async function loadTSGs() {
    try {
        const response = await fetch(`${API_URL}/teacher-subject-groups`);
        if (!response.ok) {
            console.error('Error fetching TSGs:', response.status, response.statusText);
            alert('Ошибка загрузки связей');
            return;
        }
        const tsgs = await response.json();
        console.log('TSGs data:', tsgs); // Отладка
        const table = document.getElementById('tsg-table');
        table.innerHTML = '';
        tsgs.forEach(tsg => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${tsg.id || 'N/A'}</td>
                <td>${tsg.teacher?.fullName || 'N/A'}</td>
                <td>${tsg.subject?.name || 'N/A'}</td>
                <td>${tsg.group?.name || 'N/A'}</td>
                <td>
                    <button class="edit" onclick="editTSG(${tsg.id || 0}, ${tsg.teacherID || 0}, ${tsg.subjectID || 0}, ${tsg.groupID || 0})">Редактировать</button>
                    <button class="delete" onclick="deleteTSG(${tsg.id || 0})">Удалить</button>
                </td>
            `;
            table.appendChild(row);
        });
    } catch (error) {
        console.error('Error in loadTSGs:', error);
        alert('Ошибка загрузки связей');
    }
}

// Заполнение выпадающих списков для формы связей
async function populateSelects() {
    try {
        const teacherSelect = document.getElementById('tsg-teacher');
        const subjectSelect = document.getElementById('tsg-subject');
        const groupSelect = document.getElementById('tsg-group');

        const teachers = await (await fetch(`${API_URL}/teachers`)).json();
        const subjects = await (await fetch(`${API_URL}/subjects`)).json();
        const groups = await (await fetch(`${API_URL}/groups`)).json();

        teacherSelect.innerHTML = '<option value="">Выберите преподавателя</option>' +
            teachers.map(t => `<option value="${t.id}">${t.fullName || 'N/A'}</option>`).join('');
        subjectSelect.innerHTML = '<option value="">Выберите предмет</option>' +
            subjects.map(s => `<option value="${s.id}">${s.name || 'N/A'}</option>`).join('');
        groupSelect.innerHTML = '<option value="">Выберите группу</option>' +
            groups.map(g => `<option value="${g.id}">${g.name || 'N/A'}</option>`).join('');
    } catch (error) {
        console.error('Error populating selects:', error);
        alert('Ошибка загрузки данных для формы');
    }
}

// Добавление связи
async function addTSG() {
    const teacherID = document.getElementById('tsg-teacher').value;
    const subjectID = document.getElementById('tsg-subject').value;
    const groupID = document.getElementById('tsg-group').value;

    if (!teacherID || !subjectID || !groupID) {
        alert('Выберите все поля');
        return;
    }

    try {
        await fetch(`${API_URL}/teacher-subject-groups`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ teacherID, subjectID, groupID })
        });
        hideForm('tsg');
        loadTSGs();
    } catch (error) {
        console.error('Error adding TSG:', error);
        alert('Ошибка добавления связи');
    }
}

// Редактирование связи
async function editTSG(id, currentTeacherID, currentSubjectID, currentGroupID) {
    await populateSelects();
    document.getElementById('tsg-teacher').value = currentTeacherID;
    document.getElementById('tsg-subject').value = currentSubjectID;
    document.getElementById('tsg-group').value = currentGroupID;
    showAddForm('tsg');

    const saveButton = document.getElementById('tsg-form').querySelector('button[onclick="addTSG()"]');
    saveButton.textContent = 'Обновить';
    saveButton.onclick = async () => {
        const teacherID = document.getElementById('tsg-teacher').value;
        const subjectID = document.getElementById('tsg-subject').value;
        const groupID = document.getElementById('tsg-group').value;

        if (!teacherID || !subjectID || !groupID) {
            alert('Выберите все поля');
            return;
        }

        try {
            await fetch(`${API_URL}/teacher-subject-groups/${id}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ teacherID, subjectID, groupID })
            });
            hideForm('tsg');
            loadTSGs();
            saveButton.textContent = 'Сохранить';
            saveButton.onclick = addTSG;
        } catch (error) {
            console.error('Error editing TSG:', error);
            alert('Ошибка редактирования связи');
        }
    };
}

// Удаление связи
async function deleteTSG(id) {
    if (confirm('Вы уверены, что хотите удалить связь?')) {
        try {
            await fetch(`${API_URL}/teacher-subject-groups/${id}`, {
                method: 'DELETE'
            });
            loadTSGs();
        } catch (error) {
            console.error('Error deleting TSG:', error);
            alert('Ошибка удаления связи');
        }
    }
}