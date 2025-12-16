#!/usr/bin/env bash
set -euo pipefail

IMAGE_NAME="${IMAGE_NAME:-ascii-art-web-docker}"
CONTAINER_NAME="${CONTAINER_NAME:-dockerize}"
DOCKERFILE="${DOCKERFILE:-Dockerfile}"
PORT="${PORT:-8080}"

echo "Building image '${IMAGE_NAME}' using ${DOCKERFILE}..."
docker image build -f "${DOCKERFILE}" -t "${IMAGE_NAME}" .

echo "Restarting container '${CONTAINER_NAME}' on port ${PORT}..."
docker container rm -f "${CONTAINER_NAME}" >/dev/null 2>&1 || true
docker container run -d -p "${PORT}:${PORT}" --name "${CONTAINER_NAME}" "${IMAGE_NAME}"

echo "Images (top entries):"
docker images | head

echo "Containers:"
docker ps -a | grep "${CONTAINER_NAME}" || true
