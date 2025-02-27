import "./Footer.css"; // Import CSS for styling

function Footer() {
  return (
    <footer className="footer">
      <p>&copy; {new Date().getFullYear()} Library Management System. All rights reserved.</p>
    </footer>
  );
}

export default Footer;
