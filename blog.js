let listbox = []

// let checkboxes = document.querySelectorAll(".checkbox ");

// for(let checkbox of checkboxes) {
  //     checkbox.addEventListener("click",function(){
    //         if(this.checked == true){
      //             listbox.push(this.value);
//           }else{
//               listbox = listbox.filter(i => i == this.value)
//     }
//   })
// }

let card = []
function addcard(event) {
  
  event.preventDefault()
  
  let name_project = document.getElementById("input-nama").value; 
  let sdate = document.getElementById("input-sdate").value; 
  let edate = document.getElementById("input-edate").value; 
  let desc = document.getElementById("input-desc").value; 
  let nodeJS = document.getElementById("nodejs")
  let reactJS = document.getElementById("reactjs")
  let nextJS = document.getElementById("nextjs")
  let typeScript = document.getElementById("typescript")
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
            
            
            let nodeJSImg = ''
            let reactJSImg = ''
            let nextJSImg = ''
            let typeScriptImg = ''
            
            //Pengkondisian buat masukin img icon ke variable 
            if (nodeJS.checked == true) {
                nodeJSImg = 'berkas/nodejs.png'
            }
            if (reactJS.checked == true) {
                reactJSImg = 'berkas/react.jpg'
            }
            if (nextJS.checked == true) {
                nextJSImg = 'berkas/nextjs.png'
            }
            if (typeScript.checked == true) {
                typeScriptImg = 'berkas/typescript.png'
            }
            
            
    let blog = {
    name_project,
    sdate,
    edate,
    desc,
    //listbox,
    nodeJSImg,
    reactJSImg,
    nextJSImg,
    typeScriptImg,
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
      <a href="blog-detail.html" target="_blank" class="nama">${card[i].name_project}</a>
    </hi>
    </div>
    <div class="date-blog">
    <p>durasi : ${selisihWaktu(card[i].sdate, card[i].edate)}</p>
    </div>
    <p class="desc-blog">${card[i].desc}</p>
    <div class="logjav">
    <img src="${card[i].nodeJSImg}"style="width: 20px"/>
    <img src="${card[i].reactJSImg}"style="width: 20px"/>
    <img src="${card[i].nextJSImg}"style="width: 20px"/>
    <img src="${card[i].typeScriptImg}"style="width: 20px"/>
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
  rblog()
}, 1000)

