let listbox = []



let checkboxes = document.querySelectorAll(".checkbox ");

for(let checkbox of checkboxes) {
    checkbox.addEventListener("click",function(){
        if(this.checked == true){
            listbox.push(this.value);
          }else{
              listbox = listbox.filter(i => i !== this.value)
    }
  })
}

let card = []
function addcard(event) {
  
  event.preventDefault()
  
  let name_project = document.getElementById("input-nama").value; 
  let sdate = document.getElementById("input-sdate").value; 
  let edate = document.getElementById("input-edate").value; 
  let desc = document.getElementById("input-desc").value; 
  // let nodejs = document.getElementsByClassName("checkbox1")
  // let reactjs = document.getElementsByClassName("checkbox2")
  // let nextjs = document.getElementsByClassName("checkbox3") 
  // let typescript = document.getElementsByClassName("checkbox4")
  let image = document.getElementById("input-image").files; 
  let gambar = URL.createObjectURL(image[0])

    if (name_project == "") {
      return alert('project name harus diisi !!!')
  } else if (sdate == "") {
      return alert('start date harus diisi !!!')
  } else if (edate == "") {
      return alert('end date harus diisi !!!')
  } else if (desc == "") {
      return alert('description harus diisi !!!')
  } else if (image == "") {
      return alert('image harus diisi !!!')
  }


  
  let blog = {
    name_project,
    sdate,
    edate,
    desc,
    listbox,
    gambar,
    postAt: new Date()
  }
  
  card.push(blog)
  rblog()

}

function rblog() {
  document.getElementById("content").innerHTML = ``
  for (let i = 0; i < card.length; i++) {
    document.getElementById("content").innerHTML +=
    `<div class="blog">
    <div class="blog-image">
      <img src="${card[i].gambar}">
    </div>
    <div>
    <h1  class="nama-blog">
      <a href="#" target="_blank" class="nama">${card[i].name_project}</a>
    </hi>
    </div>
    <div class="date-blog">
    <p>durasi : ${selisihWaktu(card[i].sdate, card[i].edate)}</p>
    </div>
    <p class="desc-blog">${card[i].desc}</p>
    <div class="logjav">${card[i].listbox}
    </div>
    <div style="display:flex;">
    <button type="button" class="button1" style="margin-right:5px;">edit</button>
    <button type="button" class="button1">delete</button>
    </div>
    </div>`
  }
}

function selisihWaktu(start,end) {
  let timeNow = new Date(start)
  let timeend = new Date(end)

  let distance = timeend - timeNow
  console.log(distance)

  let miliseconds = 1000

  let distanceDay = Math.floor(distance / (miliseconds * 60 * 60 * 24))
  let distanceHours = Math.floor(distance / (miliseconds * 60 * 60))
  let distanceMinutes = Math.floor(distance / (miliseconds * 60))
  let distanceSecond = Math.floor(distance / miliseconds)

  if (distanceDay > 0) {
    return `${distanceDay} Day`
  } else if (distanceHours > 0) {
    return `${distanceHours} Hours`
  } else if (distanceMinutes > 0) {
    return `${distanceMinutes} Minutes`
  } else {
    return `${distanceSecond} Second`
  }
}

setInterval(function () {
  renderBlog()
}, 1000)

