import { postJSON } from "https://cdn.jsdelivr.net/gh/crootjs/lib@0.0.1/api.js";

const urlBackend = "https://app-auth.domcloud.io/api/verify"; 
const divStatus = document.getElementById("status");
function responseData(hasil) {
    divStatus.style.display = "block";
    
    if(hasil.status === 200) {
        divStatus.className = "success";
        divStatus.innerHTML = `<strong>Berhasil!</strong> ${hasil.data.message}`;
        console.log("Data user tervalidasi:", hasil.data);
    } else {
        divStatus.innerHTML = "Gagal memverifikasi di sisi server.";
    }
}
window.tanganiKredensialGoogle = (response) => {
    console.log("Token JWT dari Google diterima!");
    
    const payload = {
        token: response.credential
    };
    postJSON(urlBackend, payload, responseData);
};