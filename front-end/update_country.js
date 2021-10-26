
Error("its an error message");

function Error(msg) {
    alert("test");
    let err = document.getElementById('error');
    console.log(err);
    err.innerHTML = "<span>any message</span>";
    err.style.display ='block';
}