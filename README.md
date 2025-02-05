# Frontend Project

This is a frontend project built with **SolidJS** and **Vite**. It includes a modern development setup with TypeScript, SolidJS Router, and Vite for fast builds and hot module replacement (HMR).

---

## Features

- **SolidJS**: A declarative JavaScript library for building user interfaces.
- **Vite**: A fast build tool and development server.
- **TypeScript**: Static type checking for better code quality.
- **SolidJS Router**: Client-side routing for single-page applications (SPAs).
- **ES Modules**: Modern JavaScript module system.

---

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Node.js**: Version 16 or higher. Download it from [nodejs.org](https://nodejs.org/).
- **npm**: Node.js package manager (comes with Node.js).

### Installation

1. **Clone the Repository**:
   Open your terminal and run the following command to clone the repository:
   ```bash
   git clone <repository-url>
   cd front
   ```
   Replace `<repository-url>` with the actual URL of your Git repository.

2. **Install Dependencies**:
   Run the following command to install all the required dependencies:
   ```bash
   npm install
   ```
   This will install all the packages listed in `package.json`, including `solid-js`, `vite`, and `vite-plugin-solid`.

3. **Verify Installation**:
   After the installation is complete, check if everything is set up correctly by running:
   ```bash
   npm run dev
   ```
   This should start the Vite development server and open the application in your default browser at `http://localhost:5173`.

---

## Running the Project

### Development Mode

To start the development server, run:
```bash
npm run dev
```
- This will start the Vite development server.
- The application will be available at `http://localhost:5173`.
- Any changes you make to the code will automatically reload the page (Hot Module Replacement).

### Production Build

To build the project for production, run:
```bash
npm run build
```
- This will generate a production-ready build in the `dist` directory.
- The build process includes:
  - TypeScript compilation.
  - Minification of JavaScript and CSS.
  - Tree-shaking to remove unused code.

### Preview the Production Build

To preview the production build locally, run:
```bash
npm run preview
```
- This will serve the production build at `http://localhost:4173`.
- Use this to test the production build before deploying it.

---

## Project Structure

Here’s an overview of the project structure:
