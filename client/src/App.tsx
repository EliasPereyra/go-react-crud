import "./App.css";
import Form from "./components/Form";
import Header from "./components/Header";
import UsersList from "./components/UsersList";

function App() {
  return (
    <>
      <Header />
      <section className="App">
        <UsersList />
        <Form />
      </section>
    </>
  );
}

export default App;
