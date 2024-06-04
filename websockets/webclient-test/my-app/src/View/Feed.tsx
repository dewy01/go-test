import { createSignal, type Component } from 'solid-js';

const WS_URL = 'ws://localhost:8080/feed';

export const Feed: Component = () => {
  const [messages, setMessages] = createSignal([] as string[]);
  const socket = new WebSocket(WS_URL);
  socket.onmessage = (event) => {
    setMessages([...messages(), event.data]);
  };
  return (
    <div class="w-screen flex flex-col gap-5 justify-center items-center">
      <a
        href="/"
        class="border-2 rounded-md px-4 hover:bg-gray-600 active:bg-gray-700 cursor-pointer mx-5"
      >
        back
      </a>
      <span>Live Feed</span>
      <div class="m-5">
        <div class="flex-col border-2 rounded-md p-2 min-w-96 min-h-80 max-h-80 overflow-y-auto">
          {messages().map((item) => (
            <div class="flex">
              <div>{item.toString()}</div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
