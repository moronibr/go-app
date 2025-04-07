$(document).ready(function () {
  $('#user-form').on('submit', function (e) {
    e.preventDefault();

    const data = {
      name: $('#name').val(),
      email: $('#email').val(),
      password: $('#password').val(),
      status: parseInt($('#status').val())
    };

    $.ajax({
      url: '/api/add-user', // Isso envia para localhost:8000/api/add-user se acessado via localhost
      type: 'POST',
      contentType: 'application/json',
      data: JSON.stringify(data),
      success: function (res) {
        alert('Usuário cadastrado com sucesso!');
        $('#user-form')[0].reset();
      },
      error: function (xhr) {
        alert('Erro ao cadastrar usuário: ' + xhr.responseText);
      }
    });
  });
});

document.addEventListener("DOMContentLoaded", () => {
  fetch('/api/users')
    .then(response => {
      if (!response.ok) {
        throw new Error('Erro ao buscar usuários');
      }
      return response.json();
    })
    .then(users => {
      const userList = document.getElementById('user-list');
      users.forEach(user => {
        const li = document.createElement('li');
        li.textContent = `${user.name} - ${user.email}`;
        userList.appendChild(li);
      });
    })
    .catch(error => {
      console.error('Erro:', error);
    });
});
