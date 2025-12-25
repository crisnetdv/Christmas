const timeAlv = new Date("Jan 1, 2026 00:00:00").getTime();

const interval = setInterval(() => {
    const actualData = new Date().getTime();
    const timeSub = timeAlv - actualData;

    if (timeSub <= 0) {
        document.getElementById("new-year-text").style.display = "block";
        document.getElementById("new-year-text").style.animation = "slide-bottom .9s ease forwards";
        return;
    }

    document.getElementById("day").textContent = Math.floor(timeSub / (1000 * 60 * 60 * 24));
    document.getElementById("hour").textContent = Math.floor((timeSub % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    document.getElementById("minute").textContent = Math.floor((timeSub % (1000 * 60 * 60)) / (1000 * 60));
    document.getElementById("second").textContent = Math.floor((timeSub % (1000 * 60)) / 1000);

}, 1000);