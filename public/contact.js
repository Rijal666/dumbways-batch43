function submitData() {
    let name = document.getElementById("nama").value
    let email = document.getElementById("mail").value
    let phone = document.getElementById("nomor").value
    let subject = document.getElementById("subjek").value
    let message = document.getElementById("pesan").value

    let emailReceiver = "rizkirizalmualim2@gmail.com"

    if (name == "") {
        return alert('nama harus diisi')
    } else if (email == "") {
        return alert('email harus diisi')
    } else if (phone == "") {
        return alert('nomor telepon harus diisi')
    } else if (subject == "") {
        return alert('subject harus diisi')
    } else if (message == "") {
        return alert('message harus diisi')
    }

    let link = document.createElement('a')
    link.href = `mailto:${emailReceiver}?subject=${subject}&body=Hello my name ${name}, ${message}, let's talk to me asap ${phone}`
    link.click()
} 