import { createSignal, onCleanup, onMount, For, createEffect } from 'solid-js';

const MessageComponent = ({ wsUrl, height = 'h-[400px]', width = 'w-full max-w-[90vw]' }) => {
  const [messages, setMessages] = createSignal<string[]>([]);
  const [connectionStatus, setConnectionStatus] = createSignal('连接中...');
  let messagesEndRef: HTMLDivElement;

  onMount(() => {
    const ws = new WebSocket(wsUrl);

    ws.onopen = () => {
      console.log('WebSocket connected');
      setConnectionStatus('已连接');
    };

    ws.onmessage = (event) => {
      const newMessage = event.data;
      setMessages((prevMessages) => [...prevMessages, newMessage]);
    };

    ws.onclose = () => {
      console.log('WebSocket disconnected');
      setConnectionStatus('连接已断开');
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
      setConnectionStatus('连接失败');
    };

    onCleanup(() => {
      ws.close();
    });
  });

  // 自动滚动到最新消息
  createEffect(() => {
    if (messages().length > 0) {
      messagesEndRef.scrollIntoView({ behavior: 'smooth' });
    }
  });

  return (
    <div class="flex flex-col items-center">
      {/* 连接状态提示 */}
      <div class="mb-4 text-sm text-gray-600">
        状态: {connectionStatus()}
      </div>

      {/* 消息窗口 */}
      <div
        class={`${width} ${height} overflow-y-auto border border-gray-300 p-4 bg-gray-50 rounded-lg shadow-sm`}
      >
        <For each={messages()}>
          {(message, index) => (
            <div class="mb-4 p-2 bg-white border border-gray-200 rounded-md break-words">
              {message}
            </div>
          )}
        </For>
        <div ref={messagesEndRef} />
      </div>
    </div>
  );
};

export default MessageComponent;