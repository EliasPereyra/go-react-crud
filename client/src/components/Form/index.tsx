import "./styles.css";

const submitHandler = async (data: {
  name: string;
  lastname: string;
  age: number;
}) => {
  const res = await fetch("/users", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    body: JSON.stringify(data),
  });

  return res.json();
};

const submitData = e => {
  submitHandler(e.target.value);
};

function Form() {
  return (
    <form method="POST">
      <label>Name:</label>
      <input name="name" type="text" />

      <label>Lastname:</label>
      <input name="lastname" type="text" />

      <label>Age:</label>
      <input name="age" type="number" />

      <button type="submit" onSubmit={submitData}>
        Send
      </button>
    </form>
  );
}

export default Form;
