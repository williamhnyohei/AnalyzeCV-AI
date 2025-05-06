# AnalyzeCV-AI

**AnalyzeCV-AI** é uma plataforma moderna para análise inteligente de currículos (CVs) utilizando inteligência artificial.  
O sistema permite o upload de arquivos PDF, processa os dados via uma fila Kafka e gera análises automáticas sobre o perfil profissional, salvando os resultados em um banco de dados.

---

## 🔥 Funcionalidades

- Upload de currículos em PDF.
- Processamento assíncrono via Kafka.
- Análise automática usando IA (supervisionada e/ou generativa).
- Avaliação e extração de informações: senioridade, habilidades, idiomas, etc.
- Armazenamento das análises em banco de dados PostgreSQL.
- Interface React para visualizar resultados de maneira amigável.

---

## ⚙️ Tecnologias

- **Backend:** Go (Golang)
- **Mensageria:** Apache Kafka
- **Banco de Dados:** PostgreSQL
- **Frontend:** React.js
- **Infraestrutura:** GitHub Actions (CI/CD), Docker, Docker Compose
- **IA:** Análise supervisionada e/ou integração com modelos generativos

---

## 🛠️ Como contribuir

Em breve adicionaremos o guia de contribuição.  
Se quiser começar, siga o projeto de tarefas no Kanban [GitHub Projects](#).

---

## 📈 Roadmap

- [x] Estrutura inicial do backend com endpoint de healthcheck.
- [ ] Configuração de ambiente Kafka + PostgreSQL.
- [ ] Upload de PDFs.
- [ ] Processamento e análise dos documentos.
- [ ] Integração com IA para sumarização e scoring.
- [ ] Desenvolvimento do frontend em React.
- [ ] Deploy automatizado.

---

## 📜 Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.
