import { createSignal, type Component } from 'solid-js';

const WS_URL = 'ws://localhost:8080/ws';

export const Chat: Component = () => {
  const [value, setValue] = createSignal('');
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
      <span>Websocket Chat</span>
      <div class="m-5">
        <div class="flex-col border-2 rounded-md p-2 min-w-4/6 min-h-80 max-h-80 overflow-y-auto">
          {messages().map((item) => (
            <div class="flex">
              <div>Received:</div>
              <div>{item.toString()}</div>
            </div>
          ))}
        </div>
        <div>
          <input
            class="text-black m-5"
            value={value()}
            onInput={(event) => {
              setValue(event.target.value);
            }}
            placeholder="Write a message"
          />
          <button
            onClick={() => {
              socket.send(value());
              setValue('');
            }}
            class="border-2 rounded-md px-4 hover:bg-gray-600 active:bg-gray-700 "
          >
            Send
          </button>
        </div>
      </div>
    </div>
  );
};
