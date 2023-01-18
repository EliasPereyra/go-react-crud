import "./App.css";

function App() {
  return (
    <div className="App">
      <h1>Hello Go, go bring me data: </h1>
      <button
        onClick={() => {
          fetch("http://localhost:3000/users").then(data => data.json());
        }}
      >
        Get data
      </button>
      <span>{}</span>
    </div>
  );
}

export default App;
