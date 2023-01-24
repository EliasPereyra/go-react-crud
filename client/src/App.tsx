import "./App.css";
import Form from "./components/Form";
import UsersList from "./components/UsersList";

function App() {
  return (
    <section className="App">
      <UsersList />

      <Form />
    </section>
  );
}

export default App;
