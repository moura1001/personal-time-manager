// reference: https://developer.mozilla.org/en-US/docs/Web/API/MutationObserver
export class CustomControls extends HTMLElement {
    constructor() {
        super();
    }

    connectedCallback() {
        const consume = this.querySelector("button#consume") as HTMLButtonElement;
        const increase = this.querySelector("button#increase") as HTMLButtonElement;

        if(!consume || !increase) {
            return;
        }
        
        this.swapOnClick(consume, "Consume", increase);
        this.swapOnClick(increase, "Increase", consume);
    }

    swapOnClick(element: HTMLButtonElement, name: string, other: HTMLButtonElement) {
        element.addEventListener("click", function () {
            if(element.innerText === name) {
                element.innerText = "Stop";
                other.disabled = true;
            } else {
                element.innerText = name;
                other.disabled = false;
            }
        });
    }
}

customElements.define('custom-controls', CustomControls);
