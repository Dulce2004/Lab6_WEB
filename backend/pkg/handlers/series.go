package handlers

import (
	"net/http"
	"series-tracker/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// GetAllSeries obtiene todas las series
// @Summary Obtener todas las series
// @Produce json
// @Success 200 {array} models.Series
// @Router /api/series [get]
func (h *Handler) GetAllSeries(c *gin.Context) {
	var series []models.Series
	if result := h.DB.Find(&series); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, series)
}

// CreateSeries crea una nueva serie
// @Summary Crear nueva serie
// @Accept json
// @Produce json
// @Param serie body models.Series true "Datos de la serie"
// @Success 201 {object} models.Series
// @Router /api/series [post]
func (h *Handler) CreateSeries(c *gin.Context) {
	var series models.Series
	if err := c.ShouldBindJSON(&series); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.DB.Create(&series); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, series)
}

// GetSeriesByID obtiene una serie por su ID
// @Summary Obtener serie por ID
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} models.Series
// @Router /api/series/{id} [get]
func (h *Handler) GetSeriesByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var series models.Series
	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(http.StatusOK, series)
}

// UpdateSeries actualiza una serie existente
// @Summary Actualizar serie
// @Accept json
// @Produce json
// @Param id path int true "ID de la serie"
// @Param serie body models.Series true "Datos actualizados"
// @Success 200 {object} models.Series
// @Router /api/series/{id} [put]
func (h *Handler) UpdateSeries(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var series models.Series
	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serie no encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&series); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&series)
	c.JSON(http.StatusOK, series)
}

// DeleteSeries elimina una serie
// @Summary Eliminar serie
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} map[string]string
// @Router /api/series/{id} [delete]
func (h *Handler) DeleteSeries(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var series models.Series
	if result := h.DB.Delete(&series, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Serie eliminada correctamente"})
}

// UpdateStatus actualiza el estado de una serie
// @Summary Actualizar estado
// @Accept json
// @Produce json
// @Param id path int true "ID de la serie"
// @Param status body map[string]string true "Nuevo estado"
// @Success 200 {object} models.Series
// @Router /api/series/{id}/status [patch]
func (h *Handler) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var updateData struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var series models.Series
	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serie no encontrada"})
		return
	}

	h.DB.Model(&series).Update("status", updateData.Status)
	c.JSON(http.StatusOK, series)
}

// IncrementEpisode incrementa el episodio actual
// @Summary Incrementar episodio
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} models.Series
// @Router /api/series/{id}/episode [patch]
func (h *Handler) IncrementEpisode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var series models.Series
	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serie no encontrada"})
		return
	}

	if series.CurrentEpisode < series.Episodes {
		series.CurrentEpisode++
		h.DB.Save(&series)
	}

	c.JSON(http.StatusOK, series)
}

// Upvote incrementa la puntuación
// @Summary Incrementar puntuación
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} models.Series
// @Router /api/series/{id}/upvote [patch]
func (h *Handler) Upvote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var series models.Series
	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serie no encontrada"})
		return
	}

	series.Score++
	h.DB.Save(&series)
	c.JSON(http.StatusOK, series)
}

// Downvote decrementa la puntuación
// @Summary Decrementar puntuación
// @Produce json
// @Param id path int true "ID de la serie"
// @Success 200 {object} models.Series
// @Router /api/series/{id}/downvote [patch]
func (h *Handler) Downvote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var series models.Series
	if result := h.DB.First(&series, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Serie no encontrada"})
		return
	}

	series.Score--
	h.DB.Save(&series)
	c.JSON(http.StatusOK, series)
}
