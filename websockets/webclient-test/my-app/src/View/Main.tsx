import type { Component } from 'solid-js';

export const Main: Component = () => {
  return (
    <div class="flex gap-2 flex-col justify-center items-center">
      <span>Connect to:</span>
      <div class="flex gap-5 justify-center, items-center">
        <a
          class="border-2 rounded-md p-2 hover:bg-gray-600 active:bg-gray-700"
          href="/chat"
        >
          <span>Chat</span>
        </a>
        <a
          class="border-2 rounded-md p-2 hover:bg-gray-600 active:bg-gray-700"
          href="/feed"
        >
          <span>Feed</span>
        </a>
      </div>
    </div>
  );
};
