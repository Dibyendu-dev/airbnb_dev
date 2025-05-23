# Airbnb Microservices Platform

This is a microservices-based platform for Airbnb-like functionality, built using Node.js, TypeScript, and Express.

## 🏗️ Project Structure

```
airbnb/
├── HotelService/     # Hotel management service
│   ├── src/         # Source code
│   ├── prisma/      # Database schema and migrations
│   └── ...
├── BookingService/  # Booking management service
│   ├── src/         # Source code
│   ├── prisma/      # Database schema and migrations
│   └── ...
└── README.md        # This file
```

## 🚀 Services

### HotelService

- Manages hotel listings and availability
- Handles hotel details, pricing, and amenities
- Provides search and filtering capabilities

### BookingService

- Handles booking reservations
- Manages booking status and updates
- Processes payment information
- Handles guest information

## 🛠️ Technology Stack

- **Language**: TypeScript
- **Framework**: Express.js
- **Database**: PostgreSQL with Prisma ORM
- **API Documentation**: Swagger/OpenAPI
- **Logging**: Winston
- **Testing**: Jest
- **Containerization**: Docker

## 🏁 Getting Started

### Prerequisites

- Node.js (v16 or higher)
- PostgreSQL
- npm or yarn
- Docker (optional)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Dibyendu-dev/airbnb.git
   cd airbnb
   ```

2. Install dependencies for each service:

   ```bash
   # Install HotelService dependencies
   cd HotelService
   npm install

   # Install BookingService dependencies
   cd ../BookingService
   npm install
   ```

3. Set up environment variables:

   - Copy `.env.example` to `.env` in each service directory
   - Update the variables according to your environment

4. Run database migrations:

   ```bash
   # In each service directory
   npx prisma migrate dev
   ```

5. Start the services:

   ```bash
   # Start HotelService
   cd HotelService
   npm run dev

   # Start BookingService
   cd ../BookingService
   npm run dev
   ```

## 📚 API Documentation

Each service has its own API documentation. Access them at:

- HotelService: `http://localhost:3000/api-docs`
- BookingService: `http://localhost:3001/api-docs`

## 🧪 Testing

Run tests for each service:

```bash
# In each service directory
npm run test
```

## 🔄 CI/CD

The project uses GitHub Actions for continuous integration and deployment:

- Automated testing
- Code quality checks
- Docker image building
- Deployment to staging/production

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👥 Authors

- Dibyendu Das - Initial work

## 🙏 Acknowledgments

- Express.js team
- Prisma team
- TypeScript team
- All contributors and maintainers
