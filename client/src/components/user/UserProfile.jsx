import { FaUser, FaEnvelope, FaPhoneAlt, FaIdBadge } from 'react-icons/fa';
import './UserProfile.css';

const UserProfile = ({ user }) => {
  return (
    <div className="user-profile">
      <h2>My Profile</h2>
      <div className="profile-card">
        <div className="profile-info">
          <div>
            <i><FaUser /></i>
            <span><strong>Name:</strong> {user.name}</span>
          </div>
          <div>
            <i><FaEnvelope /></i>
            <span><strong>Email:</strong> {user.email}</span>
          </div>
          <div>
            <i><FaPhoneAlt /></i>
            <span><strong>Contact Number:</strong> {user.contact_number}</span>
          </div>
          <div>
            <i><FaIdBadge /></i>
            <span><strong>Library ID:</strong> {user.lib_id}</span>
          </div>
        </div>
        <button className="btn">Update Profile</button>
      </div>
    </div>
  );
};

export default UserProfile;
