import { useState, useEffect } from "react";
import type { User } from "../../types";

import UserCard from "../UserCard";

import "./styles.css";

function UsersList() {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    fetch("http://localhost:45453/users")
      .then(res => {
        if (!res.ok) {
          console.error("Hey, here the error", res.status);
        }
        return res.json();
      })
      .then(data => {
        if (!data) throw Error("data not available");

        setUsers(data);
      })
      .catch(err => console.error("Error: ", err));
  }, []);

  return (
    <section>
      <ul>
        {users.map(user => (
          <UserCard key={user._id} user={user} />
        ))}
      </ul>
    </section>
  );
}

export default UsersList;
