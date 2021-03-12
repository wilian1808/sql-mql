console.log('%c Hola mundo...', 'color:orange')

const form = document.getElementById('form')
const obtener = document.getElementById('obtener')
const pantalla = document.getElementById('response')

form.addEventListener('submit', async function(e) {
  e.preventDefault()

  let data = document.getElementById('query').value
  let newData = data.split('\n').join(' ')
  
  let query = {
    data: String(newData)
  }

  console.log(query)

  await fetch('http://localhost:5555/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(query)
  })

})

obtener.addEventListener('click', async function() {
  fetch('http://localhost:5555/result')
    .then(response => response.json())
    .then(data => {
      console.log(data)
      pantalla.textContent = data.data
    })
    .catch(error => console.error(error))
})
