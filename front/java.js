
document.addEventListener("DOMContentLoaded", function() {
document.getElementById("registerButton").addEventListener("click", function(event) {
event.preventDefault();
window.location.href = "register.html";
});
});


document.getElementById('comentarioForm').addEventListener('submit', function(event) {
  event.preventDefault();

  // Obtener los valores del formulario
const nombre = document.getElementById("nombre").value;
const correo = document.getElementById("correo").value;
const contrasena = document.getElementById("contrasena").value;
const modal = document.getElementById("myModal");
const closeModalButton = document.getElementsByClassName("close")[0];
const modalMessage = document.getElementById("modal-message");

// Crear el objeto usuario
const usuario = {
name: nombre,
email: correo,
password: contrasena
};



  console.log(usuario)
  // Enviar el usuario a la API usando fetch
  fetch('http://localhost:3000/posts', {
    mode: 'no-cors',
    method: 'POST',
    headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin':'*'
      },
      body: JSON.stringify(usuario)
     
  })
  console.log("Usuario creado")
  modalMessage.textContent = 'Usuario creado exitosamente,inisiar sesion para continuar';
          modal.style.display = "flex";
  setTimeout(() => {
    console.log('Redirigiendo al usuario a la página de inicio');
    window.location.href = "inicio.html";
  }, 2600)
  closeModalButton.onclick = function() {
    modal.style.display = "none";
  }

  window.onclick = function(event) {
    if (event.target == modal) {
      modal.style.display = "none";
    };
  }
});



    