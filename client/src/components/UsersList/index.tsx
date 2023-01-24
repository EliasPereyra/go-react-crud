import { useState, useEffect } from "react";

import type { User } from "../../types";

import "./styles.css";

function UsersList() {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    fetch("/users")
      .then(res => res.json())
      .then(data => setUsers(data));
  }, []);

  return (
    <section>
      <ul>
        {users.map(user => (
          <li key={user._id}>
            <h3>{user.name}</h3>
          </li>
        ))}
      </ul>
    </section>
  );
}

export default UsersList;
