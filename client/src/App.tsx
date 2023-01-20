import "./App.css";
import Form from "./components/Form";

function App() {
  return (
    <section className="App">
      <div>
        <h1>Hello Go, go bring me data: </h1>
        <button
          onClick={() => {
            fetch("/users")
              .then(res => res.json())
              .then(data => console.log(data));
          }}
        >
          Get data
        </button>
      </div>

      <Form />
    </section>
  );
}

export default App;
