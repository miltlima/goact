package dockerfiles

const Python = `FROM python:latest

WORKDIR /app

COPY requirements.txt /app/

RUN pip install --no-cache-dir -r requirements.txt

COPY . /app

RUN groupadd -r pythonapp && useradd -r -g pythonapp pythonapp
USER pythonapp

CMD ["python", "app.py"]`
