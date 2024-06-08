document.addEventListener("DOMContentLoaded", function() {
    const loginButton = document.querySelector("button[type='submit']");
    const modal = document.getElementById("myModal");
    const closeModalButton = document.getElementsByClassName("close")[0];
    const modalMessage = document.getElementById("modal-message");

    loginButton.addEventListener("click", function(event) {
      event.preventDefault();

      const nombre = document.getElementById("nombre").value;
      const contrasena = document.getElementById("contrasena").value;

      async function checkUsers() {
        for (let id = 1; id <= 100; id++) {
          try {
            const response = await fetch(`http://localhost:3000/posts/${id}`);
            if (!response.ok) {
              continue; // skip to the next iteration if the response is not ok
            }
            const data = await response.json();
            if (data.name === nombre && data.password === contrasena) {
              console.log(`Usuario encontrado con ID ${id}. Puede iniciar sesión.`);
              window.location.href = "reservations.html";
              return; // exit the function once the user is found
            }
          } catch (error) {
            console.error('Error:', error);
          }
        }
        modalMessage.textContent = 'Usuario no encontrado. Debe registrarse.';
        modal.style.display = "flex";
        console.log('Usuario no encontrado. Debe registrarse.');
      }

      checkUsers();
    });

    closeModalButton.onclick = function() {
      modal.style.display = "none";
    }

    window.onclick = function(event) {
      if (event.target == modal) {
        modal.style.display = "none";
      }
    }
  });
