class CustomControls extends HTMLElement {
    private timer!: HTMLDivElement;
    private timerId: ReturnType<typeof setInterval> | undefined;

    constructor() {
        super();
    }

    connectedCallback() {
        const consume = this.querySelector("button#consume") as HTMLButtonElement;
        const increase = this.querySelector("button#increase") as HTMLButtonElement;
        this.timer = this.querySelector("div#timer") as HTMLDivElement;

        if(!consume || !increase) {
            return;
        }
        
        this.swapOnClick(consume, "Consume", increase);
        this.swapOnClick(increase, "Increase", consume);
    }

    swapOnClick(element: HTMLButtonElement, name: string, other: HTMLButtonElement) {
        element.addEventListener("click", () => {
            if(element.innerText === name) {
                this.startTimer(name);
                element.innerText = "Stop";
                other.disabled = true;
            } else {
                this.endTimer();
                element.innerText = name;
                other.disabled = false;
            }
        });
    }

    humanTime(ms: number) {
        const out = [`${ms % 1000}ms`];
        
        var h, m, s;
        s = Math.floor(ms / 1000);
        m = Math.floor(s / 60);
        s = s % 60;
        h = Math.floor(m / 60);
        m = m % 60;
        h = h % 24;

        if(s > 0) {
            out.push(`${s}s`)
        }
        if(m > 0) {
            out.push(`${m}m`)
        }
        if(h > 0) {
            out.push(`${h}h`)
        }

        return out.reverse().join(" ");
    }

    startTimer(name: string){
        if(name === "Increase") {
            this.timer.classList.remove("text-red-600");
            this.timer.classList.add("text-green-600");
        } else {
            this.timer.classList.add("text-red-600");
            this.timer.classList.remove("text-green-600");
        }

        const start = Date.now();

        this.timerId = setInterval(() => {
            const elapsed = Date.now() - start;
            this.timer.innerText = this.humanTime(elapsed);
        }, 16);
    }
    
    endTimer(){
        clearInterval(this.timerId);
        this.timerId = undefined;
    }
}

if(!customElements.get('custom-controls')) { customElements.define('custom-controls', CustomControls); }
