import './views/components/custom-controls'
import 'htmx.org';

declare global {
    interface Window { htmx: any; }
}

window.htmx = require('htmx.org');
