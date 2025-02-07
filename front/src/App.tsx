import MessageComponent from './components/Message.tsx';

const App = () => {
  const wsUrl = import.meta.env.WS_URL || "ws://localhost:8765/ws";
  return (
    <div class="max-w-4xl mx-auto p-5 h-screen flex flex-col items-center justify-center">
      <h1 class="text-2xl font-bold mb-4">实时消息窗口</h1>
      <MessageComponent wsUrl={wsUrl} height="h-[500px]" width="w-full max-w-[90vw]" />
    </div>
  );
};

export default App;