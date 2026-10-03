package gateway

func (r *Receiver) Stop() {
	if r.ctx != nil && r.cancel != nil {
		r.cancel()
	}
	r.logger.Info("Disconnected from Discord")
}
