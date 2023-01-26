import type { User } from "../../types";

import logo from "../../assets/user-logo.svg";
import edit from "../../assets/edit.svg";
import deleteIcon from "../../assets/delete.svg";

import "./styles.css";

function UserCard({ user }: { user: User }) {
  return (
    <figure key={user._id}>
      <figcaption>
        <img className="user-icon" src={logo} alt={user.name} />
        <div className="icons">
          <img className="edit-btn" src={edit} alt={"edit"} />
          <img className="edit-btn" src={deleteIcon} alt={"delete"} />
        </div>
      </figcaption>
      <footer>
        <h3>
          {user.name} - {user.lastname}
        </h3>
        <p>
          {user.age} <small>years old</small>
        </p>
      </footer>
    </figure>
  );
}

export default UserCard;
