import 'core-js/stable';
import 'regenerator-runtime/runtime';

import App from './App.svelte';
import './utils.css';

import Wails from '@wailsapp/runtime';

Wails.Init(() => {
    new App({
	target: document.body,
    });
});
