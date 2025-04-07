$(document).ready(function () {
    $.get("/api/users", function (data) {
      data.forEach(function (user) {
        $('#user-list').append(`<li>${user.name} - ${user.email}</li>`);
      });
    });
  });
  