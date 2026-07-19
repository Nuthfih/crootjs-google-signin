import { postJSON } from "https://cdn.jsdelivr.net/gh/crootjs/lib@0.0.1/api.js";

const urlBackend = "https://tugas-google-nuthfih.sao.dom.my.id/api/verify";
const divStatus = document.getElementById("status");
const loginArea = document.getElementById("login-area");
const userArea = document.getElementById("user-area");
const welcomeText = document.getElementById("welcome-text");
const logoutBtn = document.getElementById("logout-btn");

function responseData(hasil) {
    divStatus.style.display = "block";

    if (hasil.status === 200) {
        divStatus.className = "success";
        divStatus.innerHTML = `<strong>Berhasil!</strong> ${hasil.data.message}`;

        loginArea.style.display = "none";
        userArea.style.display = "block";
        welcomeText.innerText = `Selamat datang! ID Unikmu: ${hasil.data.user_id}`;
        console.log("Data user tervalidasi dari Golang:", hasil.data);
    } else {
        divStatus.innerHTML = "Gagal memverifikasi di sisi server. Cek console.";
        console.error("Respon server:", hasil);
    }
}
window.tanganiKredensialGoogle = (response) => {
    console.log("Token JWT dari Google diterima!");
    const payload = { token: response.credential };
    postJSON(urlBackend, payload, responseData);
};

logoutBtn.addEventListener("click", () => {
    userArea.style.display = "none";
    divStatus.style.display = "none";
    loginArea.style.display = "block";
    console.log("User telah logout.");
});

const scriptGoogle = document.createElement('script');
scriptGoogle.src = "https://accounts.google.com/gsi/client";
scriptGoogle.async = true;
scriptGoogle.defer = true;
document.head.appendChild(scriptGoogle);