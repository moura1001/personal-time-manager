import './views/components/my-component';

const app = document.getElementById('app');
const myComponent = document.createElement('my-component');

if (app) {
    app.appendChild(myComponent);
}
