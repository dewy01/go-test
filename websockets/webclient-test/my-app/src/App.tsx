import type { Component } from 'solid-js';
import { Router } from '@solidjs/router';
import { Main } from './View/Main';
import { Chat } from './View/Chat';
import { Feed } from './View/Feed';
import styles from './App.module.css';

const routes = [
  { path: '/', component: <Main /> },
  { path: '/chat', component: <Chat /> },
  { path: '/feed', component: <Feed /> },
];

const App: Component = () => {
  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <Router>{routes as any}</Router>
      </header>
    </div>
  );
};

export default App;
