import React, { useState } from "react";
import "./styles.css";

function Form() {
  const [nameInput, setNameInput] = useState("");
  const [lastNameInput, setLastNameInput] = useState("");
  const [ageInput, setAgeInput] = useState(1);

  const handleName = (e: React.FormEvent<HTMLInputElement>) => {
    setNameInput(e.currentTarget.value);
  };

  const handleLastName = (e: React.FormEvent<HTMLInputElement>) => {
    setLastNameInput(e.currentTarget.value);
  };

  const handleAge = (e: React.FormEvent<HTMLInputElement>) => {
    setAgeInput(Number(e.currentTarget.value));
  };

  const submitHandler = async (e: React.FormEvent) => {
    e.preventDefault();

    const res = await fetch("/users/:id", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: nameInput,
        lastName: lastNameInput,
        age: ageInput,
      }),
    });

    if (!res.ok) throw new Error("Hey, the request failed. What did you do?");

    return res.json();
  };

  return (
    <form method="POST" onSubmit={submitHandler}>
      <label>Name:</label>
      <input value={nameInput} onChange={handleName} name="name" type="text" />

      <label>Lastname:</label>
      <input onChange={handleLastName} name="lastname" type="text" />

      <label>Age:</label>
      <input onChange={handleAge} name="age" type="number" />

      <button type="submit">Send</button>
    </form>
  );
}

export default Form;
