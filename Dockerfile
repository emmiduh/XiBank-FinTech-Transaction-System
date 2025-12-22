FROM python:3.12-slim

WORKDIR /app

COPY requirement.txt .
RUN pip install --no-cache-dir -r requirement.txt

COPY fraud.py .

EXPOSE 5000

CMD ["python", "fraud.py"]