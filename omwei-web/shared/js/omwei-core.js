const OMWEI = {
    generatePulse: () => {
        const hex = "0123456789ABCDEF";
        let pulse = "";
        for (let i = 0; i < 32; i++) {
            pulse += hex[Math.floor(Math.random() * 16)];
            if (i % 2 === 1 && i < 31) pulse += " ";
        }
        return pulse;
    },
    init: function() {
        console.log("OM-WEI [READY]");
        const statusEl = document.querySelector('.status');
        if (statusEl) {
            setInterval(() => {
                statusEl.innerText = "PULSE: " + this.generatePulse();
            }, 1000);
        }
    }
};
window.onload = () => OMWEI.init();
