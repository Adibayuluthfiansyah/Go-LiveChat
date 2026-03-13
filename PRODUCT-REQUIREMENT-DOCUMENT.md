# Product Requirements Document (PRD): RAWVE

## 1. Meta Information
* **Product Name:** RAWVE Live Streaming Platform
* **Motto:** "Your Stream. Your Rules. Your Revenue."
* **Core Philosophy:** Kebebasan tanpa batas, dengan tanggung jawab moderasi yang jelas.
* **Platform:** Web Application (Backend-Heavy, API-First)

## 2. Problem Statement
Kreator konten saat ini terjebak dalam ekosistem monopoli (seperti Twitch atau YouTube Live) di mana potongan pendapatan platform sangat eksploitatif (mencapai 50%). Selain itu, sistem arsitektur aplikasi tersebut memakan banyak sumber daya (berat). Kreator butuh platform mandiri, ringan, dan memberikan hak penuh atas kepemilikan komunitas dan monetisasi mereka.

## 3. Product Goals & Vision
* **Creator-First Economy:** Membangun infrastruktur yang memungkinkan kreator menerima dukungan (donasi) 100% tanpa potongan.
* **Extreme Performance:** Menyediakan layanan *real-time chat* dengan latensi sangat rendah, tahan terhadap lonjakan penonton mendadak (*traffic spike*).
* **Decentralized Ready:** Membuka jalan untuk integrasi teknologi Web3 (*smart contracts* / Ethereum) di masa depan untuk transaksi donasi dan kepemilikan aset kreator yang transparan.

## 4. User Personas
1. **The Creator (Streamer):** Membutuhkan *dashboard* untuk mengatur judul siaran, memantau *chat* secara interaktif, dan melihat analitik tanpa antarmuka yang membingungkan.
2. **The Viewer (Penonton):** Membutuhkan akses instan ke *live chat* tanpa *lag*, bisa melakukan *login* dengan cepat (Google/Email), dan mendukung kreator favoritnya.

## 5. Core Features (MVP - Minimum Viable Product)
### A. Identity & Access Management (Powered by Clerk)
* Pendaftaran dan otentikasi menggunakan webhook Clerk.
* Validasi sesi melalui JWT Bearer Token di semua *protected endpoints*.

### B. Stream Control
* **Start/End Stream:** API untuk mengubah status *stream* (`is_live: boolean`).
* **Stream Metadata:** Pengaturan judul dan deskripsi *stream*.

### C. Real-Time Engagement (Live Chat)
* **WebSocket Connection:** *Full-duplex communication* menggunakan Gorilla WebSocket.
* **Chat History:** Mengambil 50 pesan terakhir secara otomatis saat penonton masuk ruangan.
* **Broadcasting:** Menyebarkan pesan dari satu pengguna ke puluhan ribu pengguna lain dalam satu ruang yang sama secara simultan menggunakan pola *goroutine* dan *channels*.

## 6. Technical Architecture & Stack
* **Backend Language:** Golang (Go 1.25+)
* **Architecture Pattern:** Clean Architecture (Domain, Usecase, Repository, Delivery)
* **API Framework:** Gin Web Framework
* **Real-time Engine:** Gorilla WebSocket
* **Database:** PostgreSQL (via GORM)
* **Frontend Compatibility Target:** Dirancang *API-first* agar sangat mudah dikonsumsi oleh kerangka kerja modern seperti Next.js atau React.

## 7. Future Roadmap (Post-MVP)
* **Web3 / Crypto Monetization:** Integrasi dengan *smart contracts* (Solidity/Hardhat) untuk memungkinkan sistem *tipping* / donasi terdesentralisasi langsung ke dompet digital (Ethereum/Polygon) milik kreator.
* **Distributed Pub/Sub:** Transisi memori WebSocket *hub* internal ke Redis Pub/Sub untuk skalabilitas horizontal (menjalankan banyak *server* Golang sekaligus).